package validator

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func Validate(input any) (errResponse []*ErrorResponse) {
	err := validate.Struct(input)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// Handle invalid validation error
			return []*ErrorResponse{
				{
					FailedField: "InvalidValidationError",
					Tag:         "invalid",
					Value:       err.Error(),
				},
			}
		}

		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			//element.FailedField = err.StructNamespace()
			element.FailedField = err.StructField() + "(" + err.Tag() + ": " + err.Param() + ")"
			element.Tag = err.Tag()
			element.Value = err.Param()
			errResponse = append(errResponse, &element)
		}
	}
	return errResponse
}
