package constant

import (
	"net/http"
)

const (
	SUCCESS = "success"
	ERROR   = "error"
)

const (
	SYSTEM = "SYSTEM"
)

const (
	// sort
	ProductSortCreatedAt   = "created_at"
	ProductSortBasePrice   = "base_price"
	ProductSortProductName = "name"

	// sort and direction
	SortDirectionAsc        = "asc"
	SortDirectionDesc       = "desc"
	ErrInvalidSortDirection = "invalid sort direction argument"
	ErrInvalidSort          = "invalid sort argument"
)

const (
	BindingParameterFailed = "failed to bind parameter"
	InvalidUUID            = "invalid uuid"
)

const (
	ProductCreateSuccess = "product created successfully"
	ProductCreateFailed  = "failed to create product"
	ProductGetSuccess    = "product fetched successfully"
	ProductGetFailed     = "failed to fetch product"
	ProductNotFound      = "product not found"
	ProductAlreadyExist  = "product already exist"
)

const (
	DbBeginTransactionFailed    = "failed to begin transaction: %v"
	DbRollbackTransactionFailed = "failed to rollback transaction: %v"
	DbCommitTransactionFailed   = "failed to commit transaction: %v"
	DataNotFound                = "data not found"
	DbReturnedMalformedData     = "database returned malformed data"
)

var (
	GenericHttpStatusMappings = map[string]int{
		BindingParameterFailed: http.StatusBadRequest,
		InvalidUUID:            http.StatusBadRequest,
	}

	ProductHttpStatusMappings = map[string]int{
		ProductCreateSuccess:        http.StatusCreated,
		ProductCreateFailed:         http.StatusInternalServerError,
		ProductGetSuccess:           http.StatusOK,
		ProductGetFailed:            http.StatusInternalServerError,
		ProductAlreadyExist:         http.StatusConflict,
		ProductNotFound:             http.StatusNotFound,
		DataNotFound:                http.StatusNotFound,
		DbBeginTransactionFailed:    http.StatusInternalServerError,
		DbRollbackTransactionFailed: http.StatusInternalServerError,
		DbCommitTransactionFailed:   http.StatusInternalServerError,
		DbReturnedMalformedData:     http.StatusInternalServerError,
	}
)

var (
	ValidProductSort   = []string{ProductSortCreatedAt, ProductSortBasePrice, ProductSortProductName}
	ValidSortDirection = []string{SortDirectionAsc, SortDirectionDesc}
)
