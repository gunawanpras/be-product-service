package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	postgres "github.com/gunawanpras/be-product-service/internal/adapter/repository/postgres/product"
	"github.com/gunawanpras/be-product-service/internal/core/product/port"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type (
	mockUUIDHelper struct {
		id uuid.UUID
	}
)

func (m mockUUIDHelper) New() uuid.UUID {
	return m.id
}

func TestNew(t *testing.T) {
	type args struct {
		attr postgres.InitAttribute
	}

	tests := []struct {
		name   string
		mockFn func(args *args, db *sqlx.DB, mock sqlmock.Sqlmock)
		want   port.Repository
	}{
		{
			name: "missing database client",
			mockFn: func(args *args, db *sqlx.DB, mock sqlmock.Sqlmock) {
				args.attr = postgres.InitAttribute{
					DB: postgres.DB{
						Db: nil,
					},
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := args{}

			db, mock, _ := sqlmock.New()
			dbx := sqlx.NewDb(db, "sqlmock")

			if tt.mockFn != nil {
				tt.mockFn(&args, dbx, mock)
			}

			assert.Panics(t, func() {
				postgres.New(args.attr)
			})
		})
	}
}
