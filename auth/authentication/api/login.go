package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"
)

// swagger:route POST /login account login
// Returns a message with StatusOK and valid JWT if successful
// responses:
// 200: successMessageWithJWT
// 400: badRequest
// 401: unauthorizedError

// Login logins a user with an email and password
func (h *Handler) Login(c echo.Context) error {
	h.logger.Info("Login method is starting to work on request...")

	// Read request body that should contain an username and a password
	var userLogin UserLoginTrial
	err := json.NewDecoder(c.Request().Body).Decode(&userLogin)
	if err != nil {
		h.logger.Error("You have to provide a body to your request with an username and password!")
		return echo.ErrBadRequest
	}

	// Connect to Postgres database
	conn := getPSQLConnection(h)
	defer conn.Close()

	// Retrieve user information from the database if user exists
	var user UserLogin
	err = conn.QueryRow("select username, passwordhash, salt from login where username=$1;", userLogin.Username).Scan(&user.username, &user.passwordHash, &user.salt)
	if err != nil {
		h.logger.Error("QueryRow failed due to:", err)
		os.Exit(1)
	}

	if user.username == "" {
		h.logger.Error("Login failed...")
		return echo.ErrUnauthorized
	}

	// Check if the password that the user provides is the same as the hashed password in the database
	passwordHash := argon2.IDKey([]byte(userLogin.Password), []byte(user.salt), 1, 64*1024, 4, 32)
	isPasswordValid := bytes.Compare(passwordHash, []byte(user.passwordHash))
	if isPasswordValid != 0 {
		h.logger.Error("Login failed...")
		return echo.ErrUnauthorized
	}

	// Create a JWT
	h.logger.Info("Creating a JWT...")
	// Private Key needs to be of size: key_size
	// The security level is: key_size / 2
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
			"name": userLogin.Username,
		})

	// Sign the JWT with the private key
	signedToken, err := token.SignedString(h.privateKey)
	if err != nil {
		h.logger.Error("Signing of JWT failed due to:", err)
		return echo.ErrBadRequest
	}

	h.logger.Info("Token was read successfully...")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	return c.JSON(http.StatusOK, map[string]string{
		"token": signedToken,
	})
}
