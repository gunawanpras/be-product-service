package response

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/be-product-service/pkg/util/constant"
	"github.com/gunawanpras/be-product-service/pkg/validator"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ValidatorResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

func OK(ctx *fiber.Ctx, message string, data any, statusMappings map[string]int) error {
	var code int

	if c, ok := statusMappings[message]; ok {
		code = c
	}

	response := Response{
		Status:  constant.SUCCESS,
		Message: message,
		Data:    data,
	}

	return ctx.Status(code).JSON(response)
}

func Error(ctx *fiber.Ctx, message string, err error, statusMappings map[string]int) error {
	var code = http.StatusNotImplemented

	if c, ok := statusMappings[err.Error()]; ok {
		code = c
	} else if c, ok := statusMappings[message]; ok {
		code = c
	}

	response := Response{
		Status:  constant.ERROR,
		Message: fmt.Sprintf("%s. reason: %v", message, err.Error()),
	}

	return ctx.Status(code).JSON(response)
}

func ErrorValidator(ctx *fiber.Ctx, err []*validator.ErrorResponse) error {
	var code = http.StatusBadRequest

	response := ValidatorResponse{
		Status:  constant.ERROR,
		Message: err,
	}

	return ctx.Status(code).JSON(response)
}
