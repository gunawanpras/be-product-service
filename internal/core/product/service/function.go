package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
	"github.com/gunawanpras/be-product-service/pkg/util/constant"
	"github.com/gunawanpras/be-product-service/pkg/util/timeutil"
)

// CreateProduct creates a new product in the system. It first checks if a product with the
// same category ID and name already exists. If not, it proceeds to create the product
// with the provided details and assigns a new ID to it.
//
// Parameters:
// - ctx: Context for controlling the lifetime of the request.
// - product: domain.Product containing the details of the product to be created.
//
// Returns:
// - res: domain.Product representing the newly created product.
// - err: error if an error occurs during the creation process.
func (service *ProductService) CreateProduct(ctx context.Context, product domain.Product) (res domain.Product, err error) {
	result, err := service.repo.ProductRepo.GetProductByName(ctx, product.CategoryID, product.Name)
	if err != nil {
		if err.Error() != constant.DataNotFound {
			return res, err
		}
	}

	if result.ID != uuid.Nil {
		return res, errors.New(constant.ProductAlreadyExist)
	}

	now := timeutil.TimeHelper.Now()
	newProduct := domain.Product{
		CategoryID:  product.CategoryID,
		SupplierID:  product.SupplierID,
		UnitID:      product.UnitID,
		Name:        product.Name,
		Description: product.Description,
		BasePrice:   product.BasePrice,
		Stock:       product.Stock,
		CreatedAt:   now,
		CreatedBy:   constant.SYSTEM,
	}

	productID, err := service.repo.ProductRepo.CreateProduct(ctx, newProduct)
	if err != nil {
		return res, err
	}

	newProduct.ID = productID

	return newProduct, nil
}

// GetListProduct retrieves a list of products filtered by product name, category type,
// and sorted by a specified field and direction. It first attempts to fetch the data
// from the cache. If the data is not found in the cache, it retrieves the data from
// the database and updates the cache with the retrieved data.
//
// Parameters:
// - ctx: Context for controlling the lifetime of the request.
// - productName: The name of the product to filter by (can be a partial match).
// - categoryType: The category type to filter the products.
// - sort: The field to sort the products by.
// - direction: The direction to sort the products (asc or desc).
//
// Returns:
// - res: domain.Products representing the list of products that match the criteria.
// - err: error if an error occurs during the retrieval process.
func (service *ProductService) GetListProduct(ctx context.Context, productName, categoryType, sort, direction string) (res domain.Products, err error) {
	// get from cache first
	cache, err := service.cache.ProductCache.GetListProductCache(ctx, productName, categoryType, sort, direction)
	if err == nil && len(cache) > 0 {
		return cache, nil
	}

	// if not found, get from database
	res, err = service.repo.ProductRepo.GetListProduct(ctx, productName, categoryType, sort, direction)
	if err != nil {
		if err.Error() != constant.DataNotFound {
			return res, err
		}
	}

	// set list product to cache
	err = service.cache.ProductCache.SetListProductCache(ctx, productName, categoryType, sort, direction, res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetProductByID retrieves a product by ID from the database.
//
// Parameters:
// - ctx: Context for controlling the lifetime of the request.
// - productID: The ID of the product to retrieve.
//
// Returns:
// - res: domain.Product representing the product with the provided ID.
// - err: error if an error occurs during the retrieval process.
func (service *ProductService) GetProductByID(ctx context.Context, productID uuid.UUID) (res domain.Product, err error) {
	res, err = service.repo.ProductRepo.GetProductByID(ctx, productID)
	if err != nil {
		if err.Error() == constant.DataNotFound {
			return res, errors.New(constant.ProductNotFound)
		}

		return res, err
	}

	return res, nil
}
