# Authentication & Authorization
In this section, we want to have a look at authentication and authorization with JSON Web Tokens (JWT). We will not focus here on TLS, this can be found in the `networking/tls` directory of this git repo. TLS encrypts messages that are send over the network but the message is not encrypted on the machine. That's why we use JWT in order to encrypt client-session information on the client machine.

## How to navigate through this repo
In the directory `jwt_authentication` you can find a minimalistic example on how to use JWT. This should be your starting point.

In the directory `authentication` you can find a more realistic example on how to use JWT, we will create a Postgres database with Docker and let the user be able to create an account with a password that will be hashed and salted and then stored in the database. The user will then be to login and obtain a JWT wich will be used to authenticate and authorize the user to perform certain actions. In order to create the database, you can run the following make target:
```bash
make start_sql_db
```
But before you do that, you can change the environment variables as you like.