package postgres

import "github.com/jmoiron/sqlx"

type ProductRepository struct {
	db        DB
	statement StatementList
}

type DB struct {
	Db *sqlx.DB
}

type StatementList struct {
	CreateProduct    *sqlx.Stmt
	ListProduct      *sqlx.Stmt
	GetProductByID   *sqlx.Stmt
	GetProductByName *sqlx.Stmt
}

type InitAttribute struct {
	DB DB
}
