-- Migration 0001 Up: Create categories table
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    description   VARCHAR(255),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by    VARCHAR(36),
    updated_at    TIMESTAMP DEFAULT NULL,
    updated_by    VARCHAR(36) DEFAULT NULL
);

CREATE UNIQUE INDEX idx_categories_name ON categories(name);
