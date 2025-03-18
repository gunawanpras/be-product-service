package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	postgres "github.com/gunawanpras/be-product-service/internal/adapter/repository/postgres/product"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
	"github.com/gunawanpras/be-product-service/pkg/util/uuidutil"
	"github.com/jmoiron/sqlx"
)

var (
	expectedQueryAddProduct = `
		INSERT INTO products (
			id, 
			category_id, 
			supplier_id, 
			unit_id, 
			name, 
			description, 
			base_price, 
			stock, 
			created_at, 
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	expectedQueryGetProduct = `
		SELECT
			p.id,
			p.category_id,
			p.supplier_id,
			p.unit_id,
			p.name,
			p.description,
			p.base_price,
			p.stock,
			p.created_at,
			p.created_by,
			p.updated_at,
			p.updated_by
		FROM products p
	`

	expectedQueryListProduct = expectedQueryGetProduct + `
		JOIN categories c on p.category_id = c.id
		WHERE 1=1
	`

	expectedQueryGetProductByID = expectedQueryGetProduct + `		
		WHERE p.id = $1
	`

	expectedQueryGetProductByName = expectedQueryGetProduct + `
		WHERE 
			p.category_id = $1 AND 
			p.name = $2
	`
)

var (
	ctx                context.Context = context.Background()
	productID                          = uuid.MustParse("e5ec5a4e-509a-4260-9d16-845032971427")
	categoryID                         = uuid.MustParse("e5ec5a4e-509a-4260-9d16-845032971429")
	supplierID                         = uuid.MustParse("e5ec5a4e-509a-4260-9d16-845032971431")
	unitID                             = uuid.MustParse("e5ec5a4e-509a-4260-9d16-845032971432")
	productName                        = "Kangkung Potong 1"
	productDescription                 = "Product description"
	productBasePrice                   = 3000
	productStock                       = 100
	productCreatedAt                   = time.Now()
	productCreatedBy                   = "SYSTEM"
	productUpdatedAt                   = time.Now()
	productUpdatedBy                   = "SYSTEM"
)

func TestProductRepository_CreateProduct(t *testing.T) {
	type fields struct {
	}

	type args struct {
		ctx     context.Context
		product domain.Product
	}

	uuidutil.UUIDHelper = mockUUIDHelper{id: productID}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFn  func(mockdb sqlmock.Sqlmock)
		wantRes uuid.UUID
		wantErr bool
	}{
		{
			name: "error when create product",
			args: args{
				ctx: ctx,
				product: domain.Product{
					CategoryID:  categoryID,
					SupplierID:  supplierID,
					UnitID:      unitID,
					Name:        productName,
					Description: &productDescription,
					BasePrice:   float64(productBasePrice),
					Stock:       productStock,
					CreatedAt:   productCreatedAt,
					CreatedBy:   productCreatedBy,
				},
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectExec(regexp.QuoteMeta(expectedQueryAddProduct)).
					WithArgs(
						productID,
						categoryID,
						supplierID,
						unitID,
						productName,
						&productDescription,
						"invalid_base_price",
						productStock,
						productCreatedAt,
						productCreatedBy,
					).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("error")))
			},
			wantRes: uuid.Nil,
			wantErr: true,
		},
		{
			name: "success create product",
			args: args{
				ctx: ctx,

				product: domain.Product{
					CategoryID:  categoryID,
					SupplierID:  supplierID,
					UnitID:      unitID,
					Name:        productName,
					Description: &productDescription,
					BasePrice:   float64(productBasePrice),
					Stock:       productStock,
					CreatedAt:   productCreatedAt,
					CreatedBy:   productCreatedBy,
				},
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectExec(regexp.QuoteMeta(expectedQueryAddProduct)).
					WithArgs(
						productID,
						categoryID,
						supplierID,
						unitID,
						productName,
						&productDescription,
						float64(productBasePrice),
						productStock,
						productCreatedAt,
						productCreatedBy,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantRes: productID,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			dbx := sqlx.NewDb(db, "sqlmock")

			mock.ExpectPrepare(regexp.QuoteMeta(expectedQueryAddProduct))

			if tt.mockFn != nil {
				tt.mockFn(mock)
			}

			repo := postgres.New(postgres.InitAttribute{
				DB: postgres.DB{
					Db: dbx,
				},
			})

			gotRes, err := repo.CreateProduct(ctx, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ProductRepository.CreateProduct() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}

}

func TestProductRepository_GetProductByID(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		mockFn  func(mockdb sqlmock.Sqlmock)
		wantRes domain.Product
		wantErr bool
	}{
		{
			name: "error when get product by id",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByID)).
					WithArgs(productID).
					WillReturnError(errors.New("error"))
			},
			wantRes: domain.Product{},
			wantErr: true,
		},
		{
			name: "error when get product not found",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByID)).
					WithArgs(productID).
					WillReturnError(sql.ErrNoRows)
			},
			wantRes: domain.Product{},
			wantErr: true,
		},
		{
			name: "error when there is malformed data",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByID)).
					WithArgs(productID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, -1, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Product{},
			wantErr: true,
		},

		{
			name: "success get product by id",
			args: args{
				ctx:       ctx,
				productID: productID,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByID)).
					WithArgs(productID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, productBasePrice, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Product{
				ID:          productID,
				CategoryID:  categoryID,
				SupplierID:  supplierID,
				UnitID:      unitID,
				Name:        productName,
				Description: &productDescription,
				BasePrice:   float64(productBasePrice),
				Stock:       productStock,
				CreatedAt:   productCreatedAt,
				CreatedBy:   productCreatedBy,
				UpdatedAt:   &productUpdatedAt,
				UpdatedBy:   &productUpdatedBy,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			dbx := sqlx.NewDb(db, "sqlmock")

			mock.ExpectPrepare(regexp.QuoteMeta(expectedQueryGetProductByID))

			if tt.mockFn != nil {
				tt.mockFn(mock)
			}

			repo := postgres.New(postgres.InitAttribute{
				DB: postgres.DB{
					Db: dbx,
				},
			})

			gotRes, err := repo.GetProductByID(ctx, tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ProductRepository.GetProductByID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestProductRepository_GetProductByName(t *testing.T) {
	type args struct {
		ctx         context.Context
		categoryID  uuid.UUID
		productName string
	}

	ctx := context.Background()

	tests := []struct {
		name    string
		args    args
		mockFn  func(mockdb sqlmock.Sqlmock)
		wantRes domain.Product
		wantErr bool
	}{
		{
			name: "error when get product by name",
			args: args{
				ctx:         ctx,
				categoryID:  categoryID,
				productName: productName,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByName)).
					WithArgs(categoryID, productName).
					WillReturnError(errors.New("error"))
			},
			wantRes: domain.Product{},
			wantErr: true,
		},
		{
			name: "error when get product not found",
			args: args{
				ctx:         ctx,
				categoryID:  categoryID,
				productName: productName,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByName)).
					WithArgs(categoryID, productName).
					WillReturnError(sql.ErrNoRows)
			},
			wantRes: domain.Product{},
			wantErr: true,
		},
		{
			name: "error when there is malformed data",
			args: args{
				ctx:         ctx,
				categoryID:  categoryID,
				productName: productName,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByName)).
					WithArgs(categoryID, productName).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, -1, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Product{},
			wantErr: true,
		},
		{
			name: "success get product by name",
			args: args{
				ctx:         ctx,
				categoryID:  categoryID,
				productName: productName,
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryGetProductByName)).
					WithArgs(categoryID, productName).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, productBasePrice, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Product{
				ID:          productID,
				CategoryID:  categoryID,
				SupplierID:  supplierID,
				UnitID:      unitID,
				Name:        productName,
				Description: &productDescription,
				BasePrice:   float64(productBasePrice),
				Stock:       productStock,
				CreatedAt:   productCreatedAt,
				CreatedBy:   productCreatedBy,
				UpdatedAt:   &productUpdatedAt,
				UpdatedBy:   &productUpdatedBy,
			},
			wantErr: false,
		},
	}

	db, mock, _ := sqlmock.New()
	dbx := sqlx.NewDb(db, "sqlmock")

	repo := postgres.New(postgres.InitAttribute{
		DB: postgres.DB{
			Db: dbx,
		},
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectPrepare(regexp.QuoteMeta(expectedQueryGetProductByName))

			if tt.mockFn != nil {
				tt.mockFn(mock)
			}

			gotRes, err := repo.GetProductByName(ctx, tt.args.categoryID, tt.args.productName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.GetProductByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ProductRepository.GetProductByName() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestProductRepository_GetListProduct(t *testing.T) {
	type args struct {
		ctx          context.Context
		productName  string
		categoryType string
		sort         string
		direction    string
	}

	ctx := context.Background()

	tests := []struct {
		name    string
		args    args
		mockFn  func(mockdb sqlmock.Sqlmock)
		wantRes domain.Products
		wantErr bool
	}{
		{
			name: "error when get list product",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "",
				sort:         "",
				direction:    "",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnError(errors.New("error"))
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "error when product not found",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "",
				sort:         "",
				direction:    "",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnError(sql.ErrNoRows)
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "error when there is malformed data",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "",
				sort:         "",
				direction:    "",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, -1, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "success get product list",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "",
				sort:         "",
				direction:    "",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, productBasePrice, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Products{
				{
					ID:          productID,
					CategoryID:  categoryID,
					SupplierID:  supplierID,
					UnitID:      unitID,
					Name:        productName,
					Description: &productDescription,
					BasePrice:   float64(productBasePrice),
					Stock:       productStock,
					CreatedAt:   productCreatedAt,
					CreatedBy:   productCreatedBy,
					UpdatedAt:   &productUpdatedAt,
					UpdatedBy:   &productUpdatedBy,
				},
			},
			wantErr: false,
		},
		{
			name: "success to filter product list by category type",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "Sayuran",
				sort:         "",
				direction:    "",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, productBasePrice, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Products{
				{
					ID:          productID,
					CategoryID:  categoryID,
					SupplierID:  supplierID,
					UnitID:      unitID,
					Name:        productName,
					Description: &productDescription,
					BasePrice:   float64(productBasePrice),
					Stock:       productStock,
					CreatedAt:   productCreatedAt,
					CreatedBy:   productCreatedBy,
					UpdatedAt:   &productUpdatedAt,
					UpdatedBy:   &productUpdatedBy,
				},
			},
			wantErr: false,
		},
		{
			name: "success to filter product list by category type, sort by base price, and direction is asc",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "sayuran",
				sort:         "base_price",
				direction:    "asc",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, productBasePrice, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Products{
				{
					ID:          productID,
					CategoryID:  categoryID,
					SupplierID:  supplierID,
					UnitID:      unitID,
					Name:        productName,
					Description: &productDescription,
					BasePrice:   float64(productBasePrice),
					Stock:       productStock,
					CreatedAt:   productCreatedAt,
					CreatedBy:   productCreatedBy,
					UpdatedAt:   &productUpdatedAt,
					UpdatedBy:   &productUpdatedBy,
				},
			},
			wantErr: false,
		},
		{
			name: "success to partial-match search product list by product name",
			args: args{
				ctx:          ctx,
				productName:  "kangkung",
				categoryType: "",
				sort:         "",
				direction:    "",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, productBasePrice, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Products{
				{
					ID:          productID,
					CategoryID:  categoryID,
					SupplierID:  supplierID,
					UnitID:      unitID,
					Name:        productName,
					Description: &productDescription,
					BasePrice:   float64(productBasePrice),
					Stock:       productStock,
					CreatedAt:   productCreatedAt,
					CreatedBy:   productCreatedBy,
					UpdatedAt:   &productUpdatedAt,
					UpdatedBy:   &productUpdatedBy,
				},
			},
			wantErr: false,
		},
		{
			name: "success to partial-match search product list by product name, filter by category type, sort by product name, and direction is desc",
			args: args{
				ctx:          ctx,
				productName:  "kangkung",
				categoryType: "sayuran",
				sort:         "name",
				direction:    "desc",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "category_id", "supplier_id", "unit_id", "name", "description", "base_price", "stock", "created_at", "created_by", "updated_at", "updated_by"}).
						AddRow(productID, categoryID, supplierID, unitID, productName, productDescription, productBasePrice, productStock, productCreatedAt, productCreatedBy, productUpdatedAt, productUpdatedBy))
			},
			wantRes: domain.Products{
				{
					ID:          productID,
					CategoryID:  categoryID,
					SupplierID:  supplierID,
					UnitID:      unitID,
					Name:        productName,
					Description: &productDescription,
					BasePrice:   float64(productBasePrice),
					Stock:       productStock,
					CreatedAt:   productCreatedAt,
					CreatedBy:   productCreatedBy,
					UpdatedAt:   &productUpdatedAt,
					UpdatedBy:   &productUpdatedBy,
				},
			},
			wantErr: false,
		},
		{
			name: "error when get list product with invalid sort",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "",
				sort:         "updated_at",
				direction:    "",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnError(errors.New("error"))
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "error when get list product with invalid direction",
			args: args{
				ctx:          ctx,
				productName:  "",
				categoryType: "",
				sort:         "",
				direction:    "invalid_direction",
			},
			mockFn: func(mockdb sqlmock.Sqlmock) {
				mockdb.ExpectQuery(regexp.QuoteMeta(expectedQueryListProduct)).
					WillReturnError(errors.New("error"))
			},
			wantRes: nil,
			wantErr: true,
		},
	}

	db, mock, _ := sqlmock.New()
	dbx := sqlx.NewDb(db, "sqlmock")

	repo := postgres.New(postgres.InitAttribute{
		DB: postgres.DB{
			Db: dbx,
		},
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn(mock)
			}

			gotRes, err := repo.GetListProduct(ctx, tt.args.productName, tt.args.categoryType, tt.args.sort, tt.args.direction)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepository.GetListProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ProductRepository.GetListProduct() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
