package handler

import (
	"fmt"
	"log"
)

func New(attr InitAttribute) *ProductHandler {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}
	return &ProductHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if !attr.Service.validate() {
		return fmt.Errorf("missing product service : %+v", attr.Service.ProductService)
	}

	return nil
}

func (service ServiceAttribute) validate() bool {
	return service.ProductService != nil
}
