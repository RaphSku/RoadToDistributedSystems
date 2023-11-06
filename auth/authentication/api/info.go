package api

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// swagger:route GET /info account info
// Returns the username when a valid JWT is being provided
// responses:
// 200: successMessageOnInfo
// 400: badRequest
// 401: unauthorizedError

// Login logins a user with an email and password
func (h *Handler) Info(c echo.Context) error {
	// The JWT is being fetched from the request's header
	token := c.Request().Header["Token"]
	if len(token) == 0 {
		h.logger.Error("No token found, please provide one!")
		return echo.ErrBadRequest
	}

	// Construct the complete JWT into one string
	tokenString := ""
	for i := range token {
		tokenString += token[i]
	}

	// Parse the JWT with custom user claims
	jwtToken, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			fmt.Println("Token Method not okay")
			return nil, echo.ErrUnauthorized
		}

		return &h.privateKey.PublicKey, nil
	})
	if err != nil {
		h.logger.Error("Parsing went wrong due to: %v", err)
		return echo.ErrUnauthorized
	}

	// Check if the JWT is valid
	if !jwtToken.Valid {
		h.logger.Error("JWT is not valid!")
		return echo.ErrUnauthorized
	}

	// Extract the custom user claims from the JWT
	claims, ok := jwtToken.Claims.(*UserClaims)
	if !ok {
		h.logger.Error("Could not extract claims due to: %v", err)
		return echo.ErrUnauthorized
	}

	username := claims.Username

	h.logger.Info("User: %v", username)

	return c.JSON(http.StatusOK, map[string]string{
		"user": username,
	})
}
