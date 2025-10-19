# Authentication & Authorization
In this section, we will look at authentication and authorization using JSON Web Tokens (JWT). We will not focus on TLS here; that is covered in the networking/tls directory of this Git repository. TLS encrypts messages transmitted over the network but does not encrypt data on the local machine. JWT is used to securely transmit and verify client session information, but it does not inherently encrypt the data — it is primarily used for signing and verifying claims.

## How to navigate through this repo
In the jwt_authentication directory, you will find a minimal example of how to use JWT. This should serve as your starting point.

In the authentication directory, you will find a more realistic example that includes creating a PostgreSQL database using Docker. In this example, users can create accounts with passwords that are hashed and salted before being stored in the database. Users can then log in and receive a JWT, which is used to authenticate and authorize them to perform certain actions. To create the database, you can run the following Make target:
```bash
make start_psql_db
```
But before you do that, create an `.env` file in this directory and add the following details:
```
POSTGRES_USER=<your_user>
POSTGRES_PASSWORD=<your_password>
POSTGRES_PORT=<port_of_your_choice>
```
