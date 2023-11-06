CREATE DATABASE auth;

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
