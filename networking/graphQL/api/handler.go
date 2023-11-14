package api

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
)

// Request Body for GraphQL requests
type postBody struct {
	Query     string `json:"query"`
	Operation string `json:"operationName,omitempty"`
}

// GraphQLHandler that handles the GraphQL endpoint
type GraphQLHandler struct {
	logger hclog.Logger
	config *GraphQLConfig
	schema graphql.Schema
}

// Create a new GraphQLHandler with the help of a GraphQLConfig
func NewGraphQLHandler(logger hclog.Logger, config *GraphQLConfig) *GraphQLHandler {
	return &GraphQLHandler{
		logger: logger,
		config: config,
	}
}

// ProductHandler responsible for the handling of the product endpoint
func (gh *GraphQLHandler) ProductHandler(c echo.Context) error {
	gh.logger.Info("Handling product request")

	// Reading the post body
	var body postBody
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		gh.logger.Info("Request body could not be decoded")
		return c.JSON(400, "Bad Request, expected body with 'query' and optionally 'operationName'")
	}

	// Check if schema was already created, if not, create it
	if gh.schema.QueryType() == nil {
		gh.schema = gh.config.createSchema()
	}

	// Evaluate the query
	result := graphql.Do(graphql.Params{
		Context:       c.Request().Context(),
		Schema:        gh.schema,
		RequestString: body.Query,
		OperationName: body.Operation,
	})

	gh.logger.Info("Finished with product request")

	return c.JSON(200, result)
}
