package service

import (
	"github.com/gunawanpras/be-product-service/config"
	"github.com/gunawanpras/be-product-service/internal/core/product/port"
)

type (
	CacheAttribute struct {
		ProductCache port.Cache
	}

	RepoAttribute struct {
		ProductRepo port.Repository
	}

	ConfigAttribute struct {
		Config *config.Config
	}

	ProductService struct {
		cache  CacheAttribute
		repo   RepoAttribute
		config ConfigAttribute
	}

	InitAttribute struct {
		Cache  CacheAttribute
		Repo   RepoAttribute
		Config ConfigAttribute
	}
)
