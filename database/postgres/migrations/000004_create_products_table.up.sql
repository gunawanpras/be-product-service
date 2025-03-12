-- Migration 0004 Up: Create products table
CREATE TABLE products (
    id    UUID PRIMARY KEY,
    category_id   UUID NOT NULL,
    supplier_id   UUID,          -- optional: product may not have a supplier
    unit_id       UUID NOT NULL, -- references units
    name          VARCHAR(150) NOT NULL,
    description   VARCHAR(255),
    base_price    NUMERIC(10,2) NOT NULL,
    stock         INTEGER DEFAULT 0,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by    VARCHAR(36),
    updated_at    TIMESTAMP DEFAULT NULL,
    updated_by    VARCHAR(36) DEFAULT NULL,
    CONSTRAINT fk_products_category FOREIGN KEY (category_id)
         REFERENCES categories(id),
    CONSTRAINT fk_products_supplier FOREIGN KEY (supplier_id)
         REFERENCES suppliers(id),
    CONSTRAINT fk_products_unit FOREIGN KEY (unit_id)
         REFERENCES units(id)
);

CREATE INDEX idx_products_category ON products(category_id);
CREATE INDEX idx_products_supplier ON products(supplier_id);
CREATE INDEX idx_products_name ON products(name);
