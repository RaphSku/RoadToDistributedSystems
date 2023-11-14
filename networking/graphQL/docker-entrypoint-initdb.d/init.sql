CREATE DATABASE api;

CREATE TABLE Products (
    id bigserial not null,
    name varchar(100) not null,
    description varchar(300) not null,
    price float4 not null,
    constraint pk_userid primary key (id)
);