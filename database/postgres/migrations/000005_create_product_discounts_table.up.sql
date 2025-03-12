-- Migration 0005 Up: Create product_discounts table
CREATE TABLE product_discounts (
    product_id         UUID PRIMARY KEY,  -- 1:1 relationship with products (if discount exists)
    discount_percent   NUMERIC(5,2),
    discount_start_date DATE,
    discount_end_date  DATE,
    max_purchase_qty   INTEGER,
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by         VARCHAR(36),
    updated_at         TIMESTAMP DEFAULT NULL,
    updated_by         VARCHAR(36) DEFAULT NULL,
    CONSTRAINT fk_pd_product FOREIGN KEY (product_id)
         REFERENCES products(id)
);
