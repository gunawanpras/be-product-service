package setup

import handler "github.com/gunawanpras/be-product-service/internal/adapter/http/handler/product"

type Handler struct {
	ProductHandler handler.Handler
}

func NewHandler(service Service) *Handler {
	return &Handler{
		ProductHandler: handler.New(handler.InitAttribute{
			Service: handler.ServiceAttribute{
				ProductService: service.ProductService,
			},
		}),
	}
}
