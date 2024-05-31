CREATE TABLE IF NOT EXISTS capsules (
                                        id SERIAL PRIMARY KEY NOT NULL,
                                        name VARCHAR(255) NOT NULL,
    user_id int
    );

create table IF NOT EXISTS stylists
(
    id                     serial,
    name                   varchar(255) not null,
    password               varchar(255) not null,
    email varchar(255) not null,
    image_url              varchar(255) not null,
    gender                 integer      not null,
    age                    integer      not null,
    experience_time        integer      not null,
    portfolio              varchar(255) not null,
    experience_description varchar(255)       not null
);

create table IF NOT EXISTS stylists_skills
(
    id         serial,
    skill      bigint not null,
    stylist_id bigint not null
);

create table IF NOT EXISTS users
(
    id        serial,
    name      varchar(255) not null,
    email varchar(255) not null,
    password  varchar      not null,
    image_url varchar(255) not null,
    gender    integer      not null,
    age       integer      not null,
    weight    integer      not null
);

create table if not exists users_login
(
    id       serial,
    name     varchar(255) not null,
    password varchar(255) not null
);

create table if not exists users_stylists
(
    id         serial,
    user_id    bigint not null,
    stylist_id bigint not null
);

CREATE TABLE IF NOT EXISTS capsules_items (
                                              id SERIAL PRIMARY KEY,
                                              capsule_id INT NOT NULL,
                                              item_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
                                     id SERIAL PRIMARY KEY NOT NULL,
                                     user_id INT NOT NULL,
                                     name VARCHAR NOT NULL,
                                     url VARCHAR,
                                     category INT NOT NULL,
                                     size_number int,
                                     size_text varchar,
                                     description varchar,
                                     color integer,
                                     file_name varchar
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY NOT NULL,
    chat_id int not null,
    user_id int not null,
    text varchar
    );

CREATE TABLE IF NOT EXISTS rooms (
    chat_id SERIAL PRIMARY KEY NOT NULL,
    user_id int not null,
    stylist_id int not null
);