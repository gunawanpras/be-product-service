package setup

import (
	"github.com/go-redis/cache/v8"
	"github.com/gunawanpras/be-product-service/config"
	setupClient "github.com/gunawanpras/be-product-service/internal/setup/client"
	"github.com/jmoiron/sqlx"
)

type ExternalServices struct {
	Postgres *sqlx.DB
	Redis    *cache.Cache
}

type CoreServices struct {
	Handler Handler
}

func InitExternalServices(conf *config.Config) *ExternalServices {
	pg := setupClient.InitPostgres(conf)
	rd := setupClient.InitRedis(conf)

	return &ExternalServices{
		Postgres: pg,
		Redis:    rd,
	}
}

func InitCoreServices(conf *config.Config, externalService *ExternalServices) *CoreServices {
	cache := NewCache(conf, externalService.Redis)
	repo := NewRepository(externalService.Postgres)
	service := NewService(conf, repo, cache)
	handler := NewHandler(service)

	return &CoreServices{
		Handler: *handler,
	}
}
