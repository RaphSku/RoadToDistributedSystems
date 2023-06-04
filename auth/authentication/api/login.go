package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	h.logger.Info("Login method is starting to work on request...")

	username := c.FormValue("username")
	password := c.FormValue("password")

	h.logger.Info("The following parameters were passed with the request:")
	h.logger.Info(username)
	h.logger.Info(password)

	if username == "Raphael" && password == "test" {
		token := jwt.New(&ECDSA)

		h.logger.Info("New token was generated...")

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Raphael Skuza"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 24)

		h.logger.Info("New claims were attributed...")

		// Private Key needs to be of size: key_size
		// The security level is: key_size / 2
		signing, _ := token.SigningString()
		h.logger.Info(signing)

		signed, _ := token.Method.Sign(signing, h.privateKey)
		h.logger.Info(signed)

		t, err := token.SignedString(h.privateKey)
		if err != nil {
			return err
		}

		h.logger.Info("Token was read successfully...")
		c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}
