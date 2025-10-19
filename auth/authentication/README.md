# How to use this code
## Requirements
- You need Go version >= 1.24.x
- Install make

## Usage
This example application uses swagger for dynamic documentation generation, so you need to install it with:
```bash
make install_dependencies
```

To see all the available make targets, run:
```bash
make
```

Before starting the server, create an `.env` file in this directory and fill out the following details:
```
DB_HOST=127.0.0.1
DB_PORT=<port_of_your_choice>
DB_NAME=auth
DB_USER=<your_user>
DB_PW=<your_password>
```

Then run the server with
```bash
make start_server
```
The server will run on localhost on port 9090.

## Documentation
A simple API documentation is provided via Swagger on the endpoint `http://localhost:9090/docs`.

## Output
In order to sign up a user, you can use the following curl command:
```bash
curl --location 'localhost:<port_of_your_choice>/signup' \
     --header 'Content-Type: application/json' \
     --data-raw '{
       "username": "raphael",
       "email": "raphael@test.com",
       "password": "test"
     }'
```
The API server will check whether the user already exists and if the email is already in use, if this is not the case, the user will be persisted to the Postgres database.

Now, we are able to login with the following curl command:
```bash
curl --location 'localhost:<port_of_your_choice>/login' \
     --header 'Content-Type: application/json' \
     --data '{
       "username": "raphael",
       "password": "test"
     }'
```
What we obtain is a JWT, in this case it looks like this:
```json
{
    "token": "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTkzOTE5ODMsIm5hbWUiOiJyYXBoYWVsIn0.nS8I2wz2Id2D7g2K5IkGsBDNBS0I-Nzor0QzRg-LR7EcEVKTROeHjTvUVZElLfezF3PS0NqVLAxOxfboLcSm_g"
}
```
With the help of this token, we can make an API call to the `/info` endpoint. Use the following curl command:
```bash
curl --location 'localhost:<port_of_your_choice>/info' \
     --header 'Token: eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTkzOTE5ODMsIm5hbWUiOiJyYXBoYWVsIn0.nS8I2wz2Id2D7g2K5IkGsBDNBS0I-Nzor0QzRg-LR7EcEVKTROeHjTvUVZElLfezF3PS0NqVLAxOxfboLcSm_g'
```
You will receive the following response:
```json 
{
    "user": "raphael"
}
```
Congratulations, with the help of JWT you are able to authenticate your users to the API server.
