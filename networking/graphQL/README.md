# GraphQL Example Project
## What is GraphQL?
GraphQL is a query language for APIs and comes with batteries included such as a runtime that handles those queries. The most prominent feature of GraphQL is that clients can specify what kind of data they want to retrieve. But GraphQL offers more like retrieving multiple resources with only one query and simplifying version control by allowing to build versionless API endpoints that are backwards compatible.

## What kind of structure do we have here?
The structure of this example project is chosen such that you can have multiple GraphQL endpoints and RestAPI endpoints next to each other. For that, create multiple directories in the `api` directory. Here, we will only use one GraphQL endpoint but the interested reader can create another endpoint and test it out.

## Make it, do it
To get an overview over all the possible make targets, just run:
```bash
make
```
For this example project, we need a Postgres database, use the following make command to start it up:
```bash
make start_db
```
After the Postgres database has started, please use the SQL statements in `docker-entrypoint-initdb.d` to setup the database and table that are needed. You can connect to your Postgres database via `psql`:
```bash
psql -h 127.0.0.1 -p 5434 -U raphael -d postgres
```
You will be prompted for the password, you can find everything related to the database in the `.env` or `postgres-compose.yaml` file. If you want to change anything, simply edit the corresponding fields in both files. 
To start the server, use:
```bash
make start_server
``` 
Now, we can start interacting with our GraphQL endpoint.

## Interacting with our GraphQL endpoint
The following curl command allows you to insert a new Product into the Postgres database:
```bash
curl --location 'http://localhost:9090/product' \
     --header 'Content-Type: application/json' \
     --data '{
        "query": "mutation Mutation{create(name:\"The Real Shoe\",description:\"Like walking over clouds\",price:199){id,name,description,price}}"
     }'
```
The output will look like this
```
{"data":{"create":{"description":"Like walking over clouds","id":1,"name":"The Real Shoe","price":199}}}
```

If you want to query only the name and price by using the id, you can use the following query:
```bash
curl --location 'http://localhost:9090/product' \
     --header 'Content-Type: application/json' \
     --data '{
        "query": "{product(id:1){id,name,price}}"
     }'
```
Since we have only inserted one product, the only id that is available is 1 and we get the following product if we query after it:
```
{"data":{"product":{"id":1,"name":"The Real Shoe","price":199}}}
```

If we want to update the price because there is a discount over the weekend, we can use the following query to do that:
```bash
curl --location 'http://localhost:9090/product' \
     --header 'Content-Type: application/json' \
     --data '{
        "query": "mutation Mutation{update(id:1,name:\"The Real Shoe\",description:\"Like walking over clouds\",price:180){id,name,description,price}}"
     }'
```
The output is the following:
```
{"data":{"update":{"description":"Like walking over clouds","id":1,"name":"The Real Shoe","price":180}}}
```

Now, let us delete this product from our database. For that, execute the following query:
```bash
curl --location 'http://localhost:9090/product' \
     --header 'Content-Type: application/json' \
     --data '{
        "query": "mutation Mutation{delete(id:1){id,name,description,price}}"       
     }'
```
The output is:
```
{"data":{"delete":{"description":"Like walking over clouds","id":1,"name":"The Real Shoe","price":170}}}
```

If you insert multiple objects into your database and afterwards want a list of all products with the id, name and price, you can use:
```bash
curl --location 'http://localhost:9090/product' \
     --header 'Content-Type: application/json' \
     --data '{
        "query": "{list{id,name,price}}" 
     }'
```
The output could look like this, depending on your products in your database:
```
{"data":{"list":[{"id":2,"name":"Mount Cycle","price":520},{"id":3,"name":"Heavy Lift","price":300.54}]}}
```
