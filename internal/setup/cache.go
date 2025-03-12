package setup

import (
	"github.com/go-redis/cache/v8"
	"github.com/gunawanpras/be-product-service/config"
	rd "github.com/gunawanpras/be-product-service/internal/adapter/cache/redis"
	productRdCache "github.com/gunawanpras/be-product-service/internal/adapter/cache/redis/product"
	"github.com/gunawanpras/be-product-service/internal/core/product/port"
)

type Cache struct {
	ProductCache port.Cache
}

func NewCache(conf *config.Config, client *cache.Cache) Cache {
	redisClient := rd.NewRedisCacheClient(rd.InitAttribute{
		Client: rd.Client{
			Client: client,
		},
	})

	productCache := productRdCache.NewProductCache(productRdCache.InitAttribute{
		RedisClient: productRdCache.RedisClient{
			RedisClient: redisClient,
		},
		Config: conf,
	})

	return Cache{
		ProductCache: productCache,
	}
}
