package product

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gunawanpras/be-product-service/internal/core/product/domain"
)

func (r *ProductCache) SetListProductCache(ctx context.Context, productName, categoryType, sort, direction string, products domain.Products) (err error) {
	cacheKey := fmt.Sprintf("products:product_name:%s:category_type:%s:sort:%s:direction:%s", productName, categoryType, sort, direction)
	cacheValue, err := json.Marshal(products)
	if err != nil {
		return err
	}
	cacheTtl := time.Duration(r.config.Redis.Primary.Ttl) * time.Minute

	err = r.redis.RedisClient.SetValue(ctx, cacheKey, cacheValue, cacheTtl)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductCache) GetListProductCache(ctx context.Context, productName, categoryType, sort, direction string) (res domain.Products, err error) {
	cacheKey := fmt.Sprintf("products:product_name:%s:category_type:%s:sort:%s:direction:%s", productName, categoryType, sort, direction)
	cacheValue, err := r.redis.RedisClient.GetValue(ctx, cacheKey)
	if err != nil {
		cacheValue = "{}"
		if !strings.Contains(err.Error(), "key is missing") {
			return res, err
		}
	}

	err = json.Unmarshal([]byte(cacheValue), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
