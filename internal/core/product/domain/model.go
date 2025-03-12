package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	CategoryID  uuid.UUID
	SupplierID  uuid.UUID
	UnitID      uuid.UUID
	Name        string
	Description *string
	BasePrice   float64
	Stock       int
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   *time.Time
	UpdatedBy   *string
}

type Products []Product
