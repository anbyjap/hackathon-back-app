CREATE DATABASE test;

\c test

CREATE TABLE Users (
    ID integer,
    Name varchar(200),
    PRIMARY KEY (ID)
);