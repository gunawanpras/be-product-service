package postgres

import (
	"fmt"
	"log"

	"github.com/gunawanpras/be-product-service/internal/core/product/port"
)

func New(attr InitAttribute) port.Repository {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	repo := &ProductRepository{
		db: attr.DB,
	}

	repo.prepareStatements()

	return repo
}

func (init InitAttribute) validate() error {
	if !init.DB.validate() {
		return fmt.Errorf("missing DB driver : %+v", init.DB)
	}

	return nil
}

func (db DB) validate() bool {
	return db.Db != nil
}
