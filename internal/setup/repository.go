package setup

import (
	productRepoPg "github.com/gunawanpras/be-product-service/internal/adapter/repository/postgres/product"
	productRepo "github.com/gunawanpras/be-product-service/internal/core/product/port"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	ProductRepo productRepo.Repository
}

func NewRepository(db *sqlx.DB) Repository {
	productRepo := productRepoPg.New(productRepoPg.InitAttribute{
		DB: productRepoPg.DB{
			Db: db,
		},
	})

	return Repository{
		ProductRepo: productRepo,
	}
}
