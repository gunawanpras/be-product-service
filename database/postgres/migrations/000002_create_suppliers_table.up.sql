-- Migration 0002 Up: Create suppliers table
CREATE TABLE suppliers (
    id UUID PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    contact_info  VARCHAR(255),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by    VARCHAR(36),
    updated_at    TIMESTAMP DEFAULT NULL,
    updated_by    VARCHAR(36) DEFAULT NULL
);

CREATE UNIQUE INDEX idx_suppliers_name ON suppliers(name);
