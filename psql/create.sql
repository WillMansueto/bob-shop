CREATE DATABASE test;

\c test

CREATE TABLE IF NOT EXISTS category(
    id SERIAL PRIMARY KEY,
    description VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS products(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price REAL NOT NULL,
    quantity INTEGER DEFAULT 0,
    amount REAL DEFAULT 0.0,
    category BIGINT NOT NULL,
    CONSTRAINT products_category_fk FOREIGN KEY(category)
    REFERENCES category(id)
);