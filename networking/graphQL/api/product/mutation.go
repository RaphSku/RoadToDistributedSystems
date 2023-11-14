package product

import (
	"fmt"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/api/graphql/utilities"
	"github.com/graphql-go/graphql"
)

// Create resolver that handles the insertion of product entities in the database
func mutationCreateResolver(params graphql.ResolveParams) (interface{}, error) {
	// Establish connection to Postgres database
	conn := utilities.GetPSQLConnection(database)
	defer conn.Close()

	// Reading out GraphQL parameters
	name := params.Args["name"].(string)
	description := params.Args["description"].(string)
	price := params.Args["price"].(float64)

	// Inserting product into Postgres database
	var product Product
	product.Name = name
	product.Description = description
	product.Price = price
	err := conn.QueryRow("insert into products(name, description, price) values ($1, $2, $3) returning id;", name, description, price).Scan(&product.ID)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed due to: %v", err)
	}

	return product, nil
}

// Update resolver that handles the updates of the products in the Postgres database
func mutationUpdateResolver(params graphql.ResolveParams) (interface{}, error) {
	// Establish connection to Postgres database
	conn := utilities.GetPSQLConnection(database)
	defer conn.Close()

	// Reading out GraphQL parameters
	id := params.Args["id"].(int)
	name := params.Args["name"].(string)
	description := params.Args["description"].(string)
	price := params.Args["price"].(float64)

	// Updating Postgres database with new product details
	var product Product
	product.ID = int64(id)
	product.Name = name
	product.Description = description
	product.Price = price
	_, err := conn.Query("update products set name=$2, description=$3, price=$4 where id=$1;", id, name, description, price)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed due to: %v", err)
	}

	return product, nil
}

// Delete resolver that handles the deletion of products in the Postgres database
func mutationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	// Establish connection to Postgres database
	conn := utilities.GetPSQLConnection(database)
	defer conn.Close()

	// Reading out GraphQL parameters
	id := params.Args["id"].(int)

	// Deleting product from Postgres database
	var product Product
	product.ID = int64(id)
	err := conn.QueryRow("delete from products where id=$1 returning name, description, price;", id).Scan(&product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed due to: %v", err)
	}

	return product, nil
}

// GraphQL object that defines the mutations that are possible on our product entity
func (po *ProductObjects) defineMutationType() {
	po.logger.Info("Mutations for product entity are created")

	po.MutationObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        po.TypeObject,
				Description: "Create new product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: mutationCreateResolver,
			},
			"update": &graphql.Field{
				Type:        po.TypeObject,
				Description: "Update product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: mutationUpdateResolver,
			},
			"delete": &graphql.Field{
				Type:        po.TypeObject,
				Description: "Delete product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: mutationDeleteResolver,
			},
		},
	})

	po.logger.Info("Mutations are defined")
}
