package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
)

type Service interface {
	CreateProduct(ctx context.Context, product domain.Product) (res domain.Product, err error)
	GetListProduct(ctx context.Context, productName, categoryType, sort, direction string) (res domain.Products, err error)
	GetProductByID(ctx context.Context, productID uuid.UUID) (res domain.Product, err error)
}
