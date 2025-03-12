package handler

import "github.com/gunawanpras/be-product-service/internal/core/product/port"

type (
	ServiceAttribute struct {
		ProductService port.Service
	}

	ProductHandler struct {
		service ServiceAttribute
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
