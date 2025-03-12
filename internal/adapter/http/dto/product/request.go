package dto

import "github.com/google/uuid"

type CreateProductRequest struct {
	CategoryID  uuid.UUID `json:"category_id" validate:"required,uuid"`
	SupplierID  uuid.UUID `json:"supplier_id" validate:"required,uuid"`
	UnitID      uuid.UUID `json:"unit_id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required,min=3,max=150"`
	Description *string   `json:"description" validate:"omitempty,max=255"`
	BasePrice   float64   `json:"base_price" validate:"required,gte=0"`
	Stock       int       `json:"stock" validate:"required,gte=0"`
}

type GetListProductRequest struct {
	ProductName string `query:"product_name" validate:"omitempty,min=3,max=150"`
	FilterSort
}

type GetProductByIDRequest struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}

type GetProductByNameRequest struct {
}

type FilterSort struct {
	CategoryType string `query:"category_type" validate:"omitempty,min=3,max=15"`
	Sort         string `query:"sort" validate:"omitempty,min=3,max=150"`
	Direction    string `query:"direction" validate:"omitempty,min=3,max=4"`
}
