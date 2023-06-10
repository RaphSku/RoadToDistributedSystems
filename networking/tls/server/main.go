package main

import (
	"log"
	"net/http"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/networking/tls/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This message is protected by TLS!")
	})

	tlsConfig, err := config.SetupTLSConfig(config.TLSConfig{
		CAFile:        config.CAFile,
		ServerAddress: "localhost:9090",
		CertFile:      config.CertFile,
		KeyFile:       config.KeyFile,
	})
	if err != nil {
		log.Fatal(err)
	}

	s := http.Server{
		Addr:      ":9090",
		Handler:   e,
		TLSConfig: tlsConfig,
	}

	if err := s.ListenAndServeTLS("./certs/server.pem", "./certs/server-key.pem"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
