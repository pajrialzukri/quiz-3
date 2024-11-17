-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(255)
);

-- +migrate StatementEnd
