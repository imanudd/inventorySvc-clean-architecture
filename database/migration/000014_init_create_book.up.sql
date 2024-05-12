-- +migrate Up
CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        author_id INT,
        book_name VARCHAR(255),
        title VARCHAR(255),
        price INT,
        created_at TIMESTAMP
);