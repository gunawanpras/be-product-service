package port

import (
	"context"

	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
)

type Cache interface {
	SetListProductCache(ctx context.Context, productName, categoryType, sort, direction string, products domain.Products) (err error)
	GetListProductCache(ctx context.Context, productName, categoryType, sort, direction string) (res domain.Products, err error)
}
