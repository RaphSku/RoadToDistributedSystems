package main

import (
	"log"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/api/graphql/api"
	"github.com/RaphSku/RoadToDistributedSystems/tree/main/api/graphql/api/product"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load the environment file
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create an application logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "graphql-product-app",
		Level: hclog.LevelFromString("INFO"),
	})

	// Create subsystem logger
	productEntityLogger := logger.Named("graphql_product")
	configProductLogger := logger.Named("graphql_product_config")
	handlerProductLogger := logger.Named("graphql_product_handler")

	// Setup the GraphQL handler responsible for the GraphQL endpoint
	productObjects := product.NewGraphQLEntity(productEntityLogger)
	productConfig := api.NewGraphQLConfig(
		configProductLogger,
		productObjects.TypeObject,
		productObjects.QueryObject,
		productObjects.MutationObject,
	)
	gh := api.NewGraphQLHandler(handlerProductLogger, productConfig)

	// Echo Server
	e := echo.New()
	e.POST("/product", gh.ProductHandler)
	e.Logger.Fatal(e.Start(":9090"))
}
