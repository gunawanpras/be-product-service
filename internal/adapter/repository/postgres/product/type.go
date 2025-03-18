package postgres

import (
	"github.com/jmoiron/sqlx"
)

type (
	ProductRepository struct {
		db        DB
		statement StatementList
	}

	DB struct {
		Db *sqlx.DB
	}

	StatementList struct {
		CreateProduct    *sqlx.Stmt
		ListProduct      *sqlx.Stmt
		GetProductByID   *sqlx.Stmt
		GetProductByName *sqlx.Stmt
	}

	InitAttribute struct {
		DB DB
	}
)
