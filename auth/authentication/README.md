# How to use this code
## Requirements
- You need Go version >= 1.20.x
- Install make

## Usage
Fetch the module dependencies with:
```bash
go mod download
```
If you want to update the documentation, you will need to install swagger, do it with:
```bash
make install_dependencies
```

Before you start the server, be sure to create a pair of keys for the JWT encryption. You can do this with the following make target:
```bash
make generate_private_public_key
```

Then run the server with
```bash
make start
```
The server will run on localhost on port 9090.

## Documentation
A simple API documentation is provided via Swagger on the endpoint `http://localhost:9090/docs`.

## Output
In order to sign up a user, you can use the following curl command:
```bash
curl --location 'localhost:9090/signup' \
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
curl --location 'localhost:9090/login' \
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
curl --location 'localhost:9090/info' \
     --header 'Token: eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTkzOTE5ODMsIm5hbWUiOiJyYXBoYWVsIn0.nS8I2wz2Id2D7g2K5IkGsBDNBS0I-Nzor0QzRg-LR7EcEVKTROeHjTvUVZElLfezF3PS0NqVLAxOxfboLcSm_g'
```
You will receive the following response:
```json 
{
    "user": "raphael"
}
```
Congratulations, with the help of JWT you are able to authenticate your users to the API server.
