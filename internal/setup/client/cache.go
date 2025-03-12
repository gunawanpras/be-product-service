package client

import (
	"fmt"
	"time"

	rCache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/gunawanpras/be-product-service/config"
)

func InitRedis(conf *config.Config) (c *rCache.Cache) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Redis.Primary.Host, conf.Redis.Primary.Port),
		DialTimeout:  time.Duration(conf.Redis.Primary.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(conf.Redis.Primary.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Redis.Primary.WriteTimeout) * time.Second,
	})

	c = rCache.New(&rCache.Options{
		Redis:      client,
		LocalCache: rCache.NewTinyLFU(1000, time.Minute),
	})

	return c
}
