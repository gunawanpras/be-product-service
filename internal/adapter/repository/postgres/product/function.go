package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
	"github.com/gunawanpras/be-product-service/pkg/util/constant"
	"github.com/gunawanpras/be-product-service/pkg/util/pageutil"
	"github.com/gunawanpras/be-product-service/pkg/util/uuidutil"
)

// CreateProduct creates a new product in the system. It assigns a new ID to the product and uses the ExecContext method of the sqlx.NamedStmt to execute the query.
//
// Parameters:
// - ctx: Context for controlling the lifetime of the request.
// - product: domain.Product containing the details of the product to be created.
//
// Returns:
// - res: uuid.UUID representing the ID of the newly created product.
// - err: error if an error occurs during the creation process.
func (repo *ProductRepository) CreateProduct(ctx context.Context, product domain.Product) (res uuid.UUID, err error) {
	product.ID = uuidutil.UUIDHelper.New()

	repo.prepareCreateProduct()
	_, err = repo.statement.CreateProduct.ExecContext(ctx, product.ID, product.CategoryID, product.SupplierID, product.UnitID, product.Name, product.Description, product.BasePrice, product.Stock, product.CreatedAt, product.CreatedBy)
	if err != nil {
		return uuid.Nil, err
	}

	return product.ID, nil
}

// GetListProduct retrieves a list of products filtered by product name, category type,
// and sorted by a specified field and direction.
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
func (repo *ProductRepository) GetListProduct(ctx context.Context, productName, categoryType, sort, direction string) (res domain.Products, err error) {
	var (
		query    []string = []string{queryGetListProduct}
		args     []any
		product  Product
		products Products
	)

	// filter by category type
	if categoryType != "" {
		query = append(query, "AND c.name = ?")
		args = append(args, categoryType)
	}

	// search by product name (partial match, case insensitive)
	if productName != "" {
		query = append(query, "AND LOWER(p.name) LIKE LOWER(?)")
		args = append(args, "%"+productName+"%")
	}

	// sort and direction
	if sort != "" {
		if err = pageutil.ValidateSortDirection(constant.ValidProductSort, sort, direction); err != nil {
			return res, err
		}

		query = append(query, "ORDER BY p."+sort+" "+direction)
	}

	finalQuery := strings.Join(query, " ")
	finalQuery = repo.db.Db.Rebind(finalQuery)

	rows, err := repo.db.Db.QueryxContext(ctx, finalQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		product = Product{}
		err = rows.StructScan(&product)
		if err != nil {
			return res, err
		}

		products = append(products, product)
	}

	if !products.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return products.ToModel(), nil
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
func (repo *ProductRepository) GetProductByID(ctx context.Context, productID uuid.UUID) (res domain.Product, err error) {
	var product Product

	repo.prepareGetProductByID()
	err = repo.statement.GetProductByID.QueryRowxContext(ctx, productID).StructScan(&product)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !product.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return product.ToModel(), nil
}

// Get product by name
func (repo *ProductRepository) GetProductByName(ctx context.Context, categoryID uuid.UUID, productName string) (res domain.Product, err error) {
	var product Product

	repo.prepareGetProductByName()
	err = repo.statement.GetProductByName.QueryRowxContext(ctx, categoryID, productName).StructScan(&product)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New(constant.DataNotFound)
		}

		return res, err
	}

	if !product.Validate() {
		return res, errors.New(constant.DbReturnedMalformedData)
	}

	return product.ToModel(), nil
}
