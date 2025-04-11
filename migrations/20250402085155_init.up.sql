CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(255)
);