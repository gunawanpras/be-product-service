package client

import (
	"log"
	"time"

	"github.com/gunawanpras/be-product-service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres(conf *config.Config) *sqlx.DB {
	pg, err := sqlx.Connect("postgres", conf.Postgre.Primary.ConnString)
	if err != nil {
		log.Panic("failed to open postgre client for product service:", err)
	}

	if err := pg.Ping(); err != nil {
		log.Panic("failed to ping PostgreSQL: %w", err)
	}

	pg.SetMaxOpenConns(conf.Postgre.Primary.MaxOpenConn)
	pg.SetMaxIdleConns(conf.Postgre.Primary.MaxIdleConn)
	pg.SetConnMaxLifetime(time.Second * time.Duration(conf.Postgre.Primary.MaxConnLifeTimeInSecond))

	return pg
}
