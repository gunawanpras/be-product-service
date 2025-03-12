package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/be-product-service/internal/setup"
)

func NewRouter(app *fiber.App, handler setup.Handler) {
	// app.Use(recover.New())
	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return nil })

	products := app.Group("/products")
	products.Post("/", handler.ProductHandler.CreateProduct)
	products.Get("/", handler.ProductHandler.GetListProduct)
	products.Get("/:id", handler.ProductHandler.GetProductByID)
}
