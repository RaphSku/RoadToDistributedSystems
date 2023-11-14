package api

import (
	"os"

	"github.com/graphql-go/graphql"
	"github.com/hashicorp/go-hclog"
)

// GraphQLConfig is a wrapper around the necessary GraphQL Objects
type GraphQLConfig struct {
	logger       hclog.Logger
	objectType   *graphql.Object
	queryType    *graphql.Object
	mutationType *graphql.Object
}

// Creates a new GraphQLConfig
func NewGraphQLConfig(logger hclog.Logger, objectType *graphql.Object, queryType *graphql.Object, mutationType *graphql.Object) *GraphQLConfig {
	return &GraphQLConfig{
		logger:       logger,
		objectType:   objectType,
		queryType:    queryType,
		mutationType: mutationType,
	}
}

// Creates a new schema based on the query object and optionally the mutation if present
func (gconfig *GraphQLConfig) createSchema() graphql.Schema {
	gconfig.logger.Info("Creating a new schema")

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    gconfig.queryType,
			Mutation: gconfig.mutationType,
		},
	)
	if err != nil {
		gconfig.logger.Error("Schema could not be created")
		os.Exit(1)
	}

	gconfig.logger.Info("Schema was created successfully")

	return schema
}
