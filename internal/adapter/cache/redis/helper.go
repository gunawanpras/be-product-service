package redis

import (
	"context"
	"time"

	rCache "github.com/go-redis/cache/v8"
)

func (r *RedisClient) SetValue(ctx context.Context, key string, value interface{}, ttl time.Duration) (err error) {
	err = r.cache.Client.Set(&rCache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   ttl,
	})

	if err != nil {
		return
	}

	return
}

func (r *RedisClient) GetValue(ctx context.Context, key string) (value string, err error) {
	if err = r.cache.Client.Get(ctx, key, &value); err != nil {
		return
	}

	return
}

func (r *RedisClient) DeleteValue(ctx context.Context, key string) (err error) {
	if err = r.cache.Client.Delete(ctx, key); err != nil {
		return
	}

	return
}
