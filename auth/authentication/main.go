//go:generate swagger generate spec
package main

import (
	"log"
	"net/http"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/auth/authentication/api"
	"github.com/RaphSku/RoadToDistributedSystems/tree/main/auth/authentication/pkg"
	"github.com/go-openapi/runtime/middleware"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	privateKeyEncoded, publicKeyEncoded := pkg.GetPrivatePublicKey("./")
	privateKey, _ := pkg.Decode(privateKeyEncoded, publicKeyEncoded)

	logger := hclog.Default()

	h := &api.Handler{}
	h.Initialize(logger, privateKey)

	e := echo.New()

	e.POST("/signup", h.SignUp)
	e.POST("/login", h.Login)
	e.GET("/info", h.Info)

	redocOptions := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	docsHandler := middleware.Redoc(redocOptions, nil)
	e.GET("/docs", echo.WrapHandler(docsHandler))
	e.GET("/swagger.yml", echo.WrapHandler(http.FileServer(http.Dir("./"))))

	e.Logger.Fatal(e.Start(":9090"))
}
