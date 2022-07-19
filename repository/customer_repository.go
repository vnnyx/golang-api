package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/vnnyx/golang-api/entity"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, tx *sqlx.Tx, customer entity.Customer) (entity.Customer, error)
	GetAllCustomer(ctx context.Context, tx *sqlx.Tx) ([]entity.Customer, error)
	GetCustomerById(ctx context.Context, tx *sqlx.Tx, customerId uint32) (entity.Customer, error)
	UpdateCustomer(ctx context.Context, tx *sqlx.Tx, customer entity.Customer) (entity.Customer, error)
	DeleteCustomer(ctx context.Context, tx *sqlx.Tx, customerId uint32) error
	GetUserByUsername(ctx context.Context, tx *sqlx.Tx, username string) (entity.Customer, error)
	DeleteAllCustomer(ctx context.Context, tx *sqlx.Tx) error
}
