package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
)

type (
	Category struct {
		CategoryID  uuid.UUID  `db:"category_id"`
		Name        string     `db:"name"`
		Description string     `db:"description"`
		CreatedAt   time.Time  `db:"created_at"`
		CreatedBy   string     `db:"created_by"`
		UpdatedAt   *time.Time `db:"updated_at"`
		UpdatedBy   *string    `db:"updated_by"`
	}

	Supplier struct {
		SupplierID  uuid.UUID  `db:"supplier_id"`
		Name        string     `db:"name"`
		ContactInfo string     `db:"contact_info"`
		CreatedAt   time.Time  `db:"created_at"`
		CreatedBy   string     `db:"created_by"`
		UpdatedAt   *time.Time `db:"updated_at"`
		UpdatedBy   *string    `db:"updated_by"`
	}

	Unit struct {
		UnitID uuid.UUID `db:"unit_id"`
		Name   string    `db:"name"`
	}

	Product struct {
		ID          uuid.UUID  `db:"id"`
		CategoryId  uuid.UUID  `db:"category_id"`
		SupplierId  uuid.UUID  `db:"supplier_id"`
		UnitId      uuid.UUID  `db:"unit_id"`
		Name        string     `db:"name"`
		Description *string    `db:"description"`
		BasePrice   float64    `db:"base_price"`
		Stock       int        `db:"stock"`
		CreatedAt   time.Time  `db:"created_at"`
		CreatedBy   string     `db:"created_by"`
		UpdatedAt   *time.Time `db:"updated_at"`
		UpdatedBy   *string    `db:"updated_by"`
	}

	ProductDiscount struct {
		ProductID         uuid.UUID  `db:"id"`
		Discount          float64    `db:"discount"`
		DiscountPercent   float64    `db:"discount_percent"`
		DiscountStartDate time.Time  `db:"discount_start_date"`
		DiscountEndDate   time.Time  `db:"discount_end_date"`
		MaxPurchaseQty    int        `db:"max_purchase_qty"`
		CreatedAt         time.Time  `db:"created_at"`
		CreatedBy         string     `db:"created_by"`
		UpdatedAt         *time.Time `db:"updated_at"`
		UpdatedBy         *string    `db:"updated_by"`
	}
)

func (p Product) Validate() bool {
	if p.ID == uuid.Nil {
		return false
	}

	if p.CategoryId == uuid.Nil {
		return false
	}

	if p.SupplierId == uuid.Nil {
		return false
	}

	if p.UnitId == uuid.Nil {
		return false
	}

	if p.Name == "" {
		return false
	}

	if p.Description != nil && *p.Description == "" {
		return false
	}

	if p.BasePrice <= 0 {
		return false
	}

	if p.Stock < 0 {
		return false
	}

	if p.CreatedAt.IsZero() {
		return false
	}

	if p.CreatedBy == "" {
		return false
	}

	if p.UpdatedAt != nil && p.UpdatedAt.IsZero() {
		return false
	}

	if p.UpdatedBy != nil && *p.UpdatedBy == "" {
		return false
	}

	return true
}

func (p Product) ToModel() domain.Product {
	return domain.Product{
		ID:          p.ID,
		CategoryID:  p.CategoryId,
		SupplierID:  p.SupplierId,
		UnitID:      p.UnitId,
		Name:        p.Name,
		Description: p.Description,
		BasePrice:   p.BasePrice,
		Stock:       p.Stock,
		CreatedAt:   p.CreatedAt,
		CreatedBy:   p.CreatedBy,
		UpdatedAt:   p.UpdatedAt,
		UpdatedBy:   p.UpdatedBy,
	}
}

type Products []Product

func (p Products) Validate() bool {
	for _, product := range p {
		if !product.Validate() {
			return false
		}
	}

	return true
}

func (p Products) ToModel() domain.Products {
	var products domain.Products

	for _, product := range p {
		products = append(products, product.ToModel())
	}

	return products
}
