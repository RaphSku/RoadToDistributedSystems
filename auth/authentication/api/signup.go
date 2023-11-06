package api

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/jackc/pgx"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"
)

// Checking is user is already registered
func checkIfUserIsRegistered(h *Handler, conn *pgx.Conn, userBody *UserSignUpBody) bool {
	var userInDB UserLogin
	row := conn.QueryRow("select username from login where username=$1;", userBody.Username)
	if row.Scan(&userInDB) != pgx.ErrNoRows {
		h.logger.Error("Username is already registered...")
		return true
	}

	return false
}

// Checking if the email that a user provides during signup is already in use or not
func checkIfEmailIsRegistered(h *Handler, conn *pgx.Conn, userBody *UserSignUpBody) bool {
	var email string
	row := conn.QueryRow("select email from login where email=$1", userBody.Email)
	if row.Scan(&email) != pgx.ErrNoRows {
		h.logger.Error("Email is already registered...")
		return true
	}

	return false
}

// Generating salt for the password hash
func generateSalt(h *Handler) []byte {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		h.logger.Error("Unable to generate a salt due to:", err)
		os.Exit(1)
	}

	return salt
}

// swagger:route POST /signup account signup
// Returns a message with StatusOK if successful
// responses:
// 200: successMessage
// 400: badRequest
// 500: internalServerError

// SignUp Signs up a user with an email and password
func (h *Handler) SignUp(c echo.Context) error {
	// Read the request body that should have an email, a username and a password
	var userBody UserSignUpBody
	err := json.NewDecoder(c.Request().Body).Decode(&userBody)
	if err != nil {
		h.logger.Error("You have to provide a body to your request with an email, username and password!")
		return echo.ErrBadRequest
	}

	// Get connection to Postgres database
	conn := getPSQLConnection(h)
	defer conn.Close()

	// Check if the user is already registered
	isRegistered := checkIfUserIsRegistered(h, conn, &userBody)
	if isRegistered {
		return echo.ErrBadRequest
	}

	// Check if email is already in use
	emailIsRegistered := checkIfEmailIsRegistered(h, conn, &userBody)
	if emailIsRegistered {
		return echo.ErrBadRequest
	}

	// Hash the password with a salt with the Argon2 algorithm
	salt := generateSalt(h)
	passwordHash := argon2.IDKey([]byte(userBody.Password), salt, 1, 64*1024, 4, 32)

	// Create a user in the database
	var user UserSignUp
	sqlStatement := "insert into login(email, username, passwordhash, salt) values ($1, $2, $3, $4) returning username, passwordhash;"
	err = conn.QueryRow(sqlStatement, userBody.Email, userBody.Username, passwordHash, salt).Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		h.logger.Error("QueryRow failed due to:", err)
		return c.String(http.StatusInternalServerError, "Request failed, please retry later...")
	}

	return c.String(http.StatusOK, "Congratulations, you are signed up")
}
