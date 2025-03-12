package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
)

type Repository interface {
	CreateProduct(ctx context.Context, product domain.Product) (res uuid.UUID, err error)
	GetListProduct(ctx context.Context, productName, categoryType, sort, direction string) (res domain.Products, err error)
	GetProductByID(ctx context.Context, productID uuid.UUID) (res domain.Product, err error)
	GetProductByName(ctx context.Context, categoryID uuid.UUID, productName string) (res domain.Product, err error)
}
