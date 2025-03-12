package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *ProductRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *ProductRepository) prepareCreateProduct() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCreateProduct); err != nil {
		log.Panic("[prepareCreateProduct] error:", err)
	}
	repo.statement.CreateProduct = stmt
}

func (repo *ProductRepository) prepareListProduct() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryListProduct); err != nil {
		log.Panic("[prepareListProduct] error:", err)
	}
	repo.statement.ListProduct = stmt
}

func (repo *ProductRepository) prepareGetProductByID() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetProductByID); err != nil {
		log.Panic("[prepareGetProductByID] error:", err)
	}
	repo.statement.GetProductByID = stmt
}

func (repo *ProductRepository) prepareGetProductByName() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetProductByName); err != nil {
		log.Panic("[prepareGetProductByName] error:", err)
	}
	repo.statement.GetProductByName = stmt
}
