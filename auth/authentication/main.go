package main

import (
	"github.com/RaphSku/RoadToDistributedSystems/tree/main/auth/authentication/api"
	"github.com/RaphSku/RoadToDistributedSystems/tree/main/auth/authentication/pkg"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	privateKeyEncoded, publicKeyEncoded := pkg.GetPrivatePublicKey("./")
	privateKey, _ := pkg.Decode(privateKeyEncoded, publicKeyEncoded)

	logger := hclog.Default()

	h := &api.Handler{}
	h.Initialize(logger, privateKey)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/login", h.Login)

	e.Logger.Fatal(e.Start(":9090"))

}
