-- Migration 0003 Up: Create units table
CREATE TABLE units (
    id   UUID PRIMARY KEY,
    unit_name VARCHAR(50) NOT NULL
);

CREATE UNIQUE INDEX idx_units_name ON units(unit_name);
