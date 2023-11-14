package product

import (
	"os"

	"github.com/graphql-go/graphql"
	"github.com/hashicorp/go-hclog"
)

// Database that will be used by the product entity
var database string

// Product structure
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// GraphQL objects related to the Product entity
type ProductObjects struct {
	logger         hclog.Logger
	TypeObject     *graphql.Object
	MutationObject *graphql.Object
	QueryObject    *graphql.Object
}

// Create a new GraphQL product entity
func NewGraphQLEntity(logger hclog.Logger) *ProductObjects {
	database = os.Getenv("PRODUCT_DATABASE")

	var productObjects ProductObjects
	productObjects.logger = logger
	productObjects.defineProductType()
	productObjects.defineMutationType()
	productObjects.defineQueryType()

	return &productObjects
}

// GraphQL object that defines the fields of our product entity
func (po *ProductObjects) defineProductType() {
	po.logger.Info("A new product entity definition is being created")

	// Create new GraphQL object that represents our product entity
	po.TypeObject = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Product",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Float,
				},
			},
		},
	)

	po.logger.Info("Product entity definition is finished")
}
