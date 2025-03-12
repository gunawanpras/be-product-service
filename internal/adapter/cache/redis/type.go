package redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
)

type (
	Cache interface {
		SetValue(ctx context.Context, key string, value any, ttl time.Duration) (err error)
		GetValue(ctx context.Context, key string) (string, error)
		DeleteValue(ctx context.Context, key string) error
	}

	RedisClient struct {
		cache Client
	}

	InitAttribute struct {
		Client Client
	}

	Client struct {
		Client *cache.Cache
	}
)
