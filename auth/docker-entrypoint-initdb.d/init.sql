-- Create a database called auth
CREATE DATABASE auth;

-- Connect to database auth
\c auth;

-- Create a table called login that will be used in the authentication example
CREATE TABLE login (
    userid bigserial not null,
    email varchar(200) not null,
    username varchar(200) not null,
    passwordhash bytea not null,
    salt bytea not null,
    unique (email),
    unique (username),
    constraint pk_userid primary key (userid)
);
