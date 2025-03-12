package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateProduct(c *fiber.Ctx) error
	GetListProduct(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
}
