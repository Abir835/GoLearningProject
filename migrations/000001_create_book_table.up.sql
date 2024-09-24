-- +migrate Up
CREATE TABLE
    IF NOT EXISTS book (
                                     id SERIAL PRIMARY KEY NOT NULL,
                                     name varchar(255) NOT NULL
    );