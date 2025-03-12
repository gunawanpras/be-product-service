package main

import (
	"github.com/gunawanpras/be-product-service/config"
	"github.com/gunawanpras/be-product-service/delivery/server"
	"github.com/gunawanpras/be-product-service/internal/setup"
)

func main() {
	// init config
	config.Init(config.WithConfigFile("config"), config.WithConfigType("yaml"))
	conf := config.Get()

	// init external services
	externalService := setup.InitExternalServices(conf)
	defer externalService.Postgres.Close()

	// init core services
	coreService := setup.InitCoreServices(conf, externalService)

	// init server
	server.Up(coreService.Handler, conf.Server)
}
