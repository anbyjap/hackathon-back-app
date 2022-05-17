CREATE DATABASE test;

\c test

CREATE TABLE users (
    id SERIAL,
    name VARCHAR(200),
    height REAL,
    sex INTEGER,
    old INTEGER,
    pass VARCHAR(200),
    state INTEGER DEFAULT 0,
    created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (id)
);

INSERT INTO users(name, height, sex, old, pass) VALUES ('Aizu Taro', 164.2, 0, 21, 'test'),
    ('Aizu Hanako', 164.2, 0, 21, 'test');

CREATE TABLE weights (
    id SERIAL,
    user_id SERIAL NOT NULL,
    value REAL,
    created_at timestamp NOT NULL default current_timestamp,
    PRIMARY KEY (id),
    CONSTRAINT fk_Users FOREIGN KEY(user_id) REFERENCES Users(id) ON DELETE CASCADE
);

INSERT INTO weights(user_id, value) VALUES (1, 78.2),
    (2, 58.8),
    (2, 45.8);