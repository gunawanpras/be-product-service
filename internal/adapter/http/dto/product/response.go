package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
)

type (
	CreateProductResponse struct {
		ID uuid.UUID `json:"id"`
	}

	GetProductResponse struct {
		ID          uuid.UUID `json:"id"`
		CategoryID  uuid.UUID `json:"category_id"`
		SupplierID  uuid.UUID `json:"supplier_id"`
		UnitID      uuid.UUID `json:"unit_id"`
		Name        string    `json:"name"`
		Description *string   `json:"description"`
		BasePrice   float64   `json:"base_price"`
		Stock       int       `json:"stock"`
		CreatedAt   string    `json:"created_at"`
		CreatedBy   string    `json:"created_by"`
	}

	GetListProductResponse []GetProductResponse
)

func (p *GetProductResponse) ToResponse(product domain.Product) {
	*p = GetProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		SupplierID:  product.SupplierID,
		UnitID:      product.UnitID,
		Name:        product.Name,
		Description: product.Description,
		BasePrice:   product.BasePrice,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		CreatedBy:   product.CreatedBy,
	}
}

func (p *GetListProductResponse) ToResponse(products domain.Products) {
	for _, product := range products {
		*p = append(*p, GetProductResponse{
			ID:          product.ID,
			CategoryID:  product.CategoryID,
			SupplierID:  product.SupplierID,
			UnitID:      product.UnitID,
			Name:        product.Name,
			Description: product.Description,
			BasePrice:   product.BasePrice,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.Format(time.RFC3339),
			CreatedBy:   product.CreatedBy,
		})
	}
}
