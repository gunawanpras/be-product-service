package product

import (
	"github.com/gunawanpras/be-product-service/config"
	"github.com/gunawanpras/be-product-service/internal/adapter/cache/redis"
)

type (
	RedisClient struct {
		RedisClient redis.Cache
	}

	ProductCache struct {
		redis  RedisClient
		config *config.Config
	}

	InitAttribute struct {
		RedisClient RedisClient
		Config      *config.Config
	}
)
