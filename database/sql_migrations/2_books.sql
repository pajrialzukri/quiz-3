-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1024),
    image_url VARCHAR(1024),
    release_year INT,
    price INT,
    total_page INT,
    thickness VARCHAR(255),
    category_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(255),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- +migrate StatementEnd
