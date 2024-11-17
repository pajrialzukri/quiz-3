-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(255)
);

-- +migrate StatementEnd
