CREATE TABLE users (
    id   STRING PRIMARY KEY,
    reservation_date TIMESTAMP UNIQUE NOT NULL,
    password  TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


