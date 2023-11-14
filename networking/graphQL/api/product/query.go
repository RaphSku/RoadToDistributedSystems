package product

import (
	"fmt"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/api/graphql/utilities"
	"github.com/graphql-go/graphql"
	"github.com/labstack/gommon/log"
)

// Resolves queries that look for a product by id
func graphqlQueryResolver(params graphql.ResolveParams) (interface{}, error) {
	var product Product
	id, ok := params.Args["id"].(int)
	if ok {
		// Establish a connection to the Postgres database
		conn := utilities.GetPSQLConnection(database)
		defer conn.Close()

		// Get the product by id
		err := conn.QueryRow("select id, name, description, price from products where id=$1;", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, fmt.Errorf("QueryRow failed due to: %v", err)
		}
	}

	return product, nil
}

// Get all products
func graphqlProductList(params graphql.ResolveParams) (interface{}, error) {
	// Establish a connection with the Postgres database
	conn := utilities.GetPSQLConnection(database)
	defer conn.Close()

	// Get all products
	rows, err := conn.Query("select id, name, description, price from products;")
	if err != nil {
		log.Error("QueryRow failed due to:", err)
		return nil, fmt.Errorf("QueryRow failed due to: %v", err)
	}
	defer rows.Close()

	// Loop over all products and store them in an array
	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, fmt.Errorf("row could not be scanned due to: %v", err)
		}

		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to load a row due to: %v", err)
	}

	return products, nil
}

// GraphQL object that defines the queries that are possible on our product entity
func (po *ProductObjects) defineQueryType() {
	po.logger.Info("Query types for product entity are created")

	po.QueryObject = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"product": &graphql.Field{
					Type:        po.TypeObject,
					Description: "Get product by id",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: graphqlQueryResolver,
				},
				"list": &graphql.Field{
					Type:        graphql.NewList(po.TypeObject),
					Description: "Get product list",
					Resolve:     graphqlProductList,
				},
			},
		},
	)

	po.logger.Info("Query types are defined")
}
