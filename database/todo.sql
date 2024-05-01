-- Active: 1709977213404@@127.0.0.1@3306@go_rest_api

SHOW TABLES;

CREATE TABLE todo (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    todo TEXT NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);


CREATE TABLE user(
    user_id VARCHAR(100) NOT NULL PRIMARY KEY,
    username VARCHAR(100) NOT NULL
);


