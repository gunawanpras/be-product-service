package handler

import (
	"github.com/gofiber/fiber/v2"
	dto "github.com/gunawanpras/be-product-service/internal/adapter/http/dto/product"
	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
	"github.com/gunawanpras/be-product-service/pkg/response"
	"github.com/gunawanpras/be-product-service/pkg/util/constant"
	"github.com/gunawanpras/be-product-service/pkg/validator"
)

// CreateProduct handles the creation of a new product. It parses the request body to
// extract product details, validates the parsed data, and then calls the ProductService
// to create the product in the system. On success, it returns a response with the ID of
// the newly created product.
//
// Parameters:
//   - c: *fiber.Ctx, the Fiber context that provides request and response handling.
//
// Returns:
//   - error: an error if any issue occurs during request parsing, validation, or product
//     creation, otherwise nil.
func (handler *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req dto.CreateProductRequest

	ctx := c.UserContext()
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	args := domain.Product{
		CategoryID:  req.CategoryID,
		SupplierID:  req.SupplierID,
		UnitID:      req.UnitID,
		Name:        req.Name,
		Description: req.Description,
		BasePrice:   req.BasePrice,
		Stock:       req.Stock,
	}

	resp, err := handler.service.ProductService.CreateProduct(ctx, args)
	if err != nil {
		return response.Error(c, constant.ProductCreateFailed, err, constant.ProductHttpStatusMappings)
	}

	respData := dto.CreateProductResponse{
		ID: resp.ID,
	}

	return response.OK(c, constant.ProductCreateSuccess, respData, constant.ProductHttpStatusMappings)
}

// GetListProduct retrieves a list of products filtered by product name, category type,
// and sorted by a specified field and direction.
//
// Parameters:
//   - c: *fiber.Ctx, the Fiber context that provides request and response handling.
//
// Returns:
//   - error: an error if any issue occurs during request parsing, validation, or product
//     retrieval, otherwise nil.
func (handler *ProductHandler) GetListProduct(c *fiber.Ctx) error {
	var (
		req dto.GetListProductRequest
		res dto.GetListProductResponse
	)

	ctx := c.UserContext()
	if err := c.QueryParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.ProductService.GetListProduct(ctx, req.ProductName, req.CategoryType, req.Sort, req.Direction)
	if err != nil {
		return response.Error(c, constant.ProductGetFailed, err, constant.ProductHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.ProductGetSuccess, res, constant.ProductHttpStatusMappings)
}

// GetProductByID retrieves a product by its unique identifier. It extracts the product ID
// from the URI, validates it, and then calls the ProductService to fetch the product details.
// On success, it returns the product information in the response.
//
// Parameters:
//   - c: *fiber.Ctx, the Fiber context that provides request and response handling.
//
// Returns:
//   - error: an error if any issue occurs during parameter parsing, validation, or product
//     retrieval, otherwise nil.
func (handler *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	var (
		req dto.GetProductByIDRequest
		res dto.GetProductResponse
	)

	ctx := c.UserContext()
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.ProductService.GetProductByID(ctx, req.ID)
	if err != nil {
		return response.Error(c, constant.ProductGetFailed, err, constant.ProductHttpStatusMappings)
	}

	res.ToResponse(resp)

	return response.OK(c, constant.ProductGetSuccess, res, constant.ProductHttpStatusMappings)
}
