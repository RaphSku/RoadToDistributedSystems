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

CREATE TABLE login_activity (
    activityid bigserial not null,
    userid bigint not null,
    username varchar(200) not null,
    timestamp timestamp not null,
    constraint pk_activityid primary key (activityid),
    constraint fk_userid foreign key (userid) references login (userid)
);