package setup

import (
	"github.com/gunawanpras/be-product-service/config"
	productPort "github.com/gunawanpras/be-product-service/internal/core/product/port"
	productService "github.com/gunawanpras/be-product-service/internal/core/product/service"
)

type Service struct {
	ProductService productPort.Service
}

func NewService(conf *config.Config, repo Repository, cache Cache) Service {
	return Service{
		ProductService: productService.New(productService.InitAttribute{
			Repo: productService.RepoAttribute{
				ProductRepo: repo.ProductRepo,
			},
			Cache: productService.CacheAttribute{
				ProductCache: cache.ProductCache,
			},
			Config: productService.ConfigAttribute{
				Config: conf,
			},
		}),
	}
}
