CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username VARCHAR(255),
    name VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);