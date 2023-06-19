package repository

import (
	"context"
	"database/sql"
	"insectdie/go-restapi-mux/model/domain"
)

type CustomerRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
}
