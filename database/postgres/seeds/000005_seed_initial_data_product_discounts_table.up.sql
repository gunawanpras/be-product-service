INSERT INTO product_discounts
    (product_id, discount_percent, discount_start_date, discount_end_date, max_purchase_qty, created_at, created_by, updated_at, updated_by)
VALUES
    ('00000000-0000-0000-0000-000000000031', 15.00, CURRENT_DATE, CURRENT_DATE + INTERVAL '7 days', 5, CURRENT_TIMESTAMP, 'SYSTEM', NULL, NULL),
    ('00000000-0000-0000-0000-000000000033', 25.00, CURRENT_DATE, CURRENT_DATE + INTERVAL '7 days', 3, CURRENT_TIMESTAMP, 'SYSTEM', NULL, NULL);
