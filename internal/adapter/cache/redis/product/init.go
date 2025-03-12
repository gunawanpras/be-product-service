package product

import (
	"fmt"
	"log"

	"github.com/gunawanpras/be-product-service/internal/core/product/port"
)

func NewProductCache(attr InitAttribute) port.Cache {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	rd := &ProductCache{
		redis:  attr.RedisClient,
		config: attr.Config,
	}

	return rd
}

func (init InitAttribute) validate() error {
	if !init.RedisClient.validate() {
		return fmt.Errorf("missing redis client : %+v", init.RedisClient)
	}

	return nil
}

func (client RedisClient) validate() bool {
	return client.RedisClient != nil
}
