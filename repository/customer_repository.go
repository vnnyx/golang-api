package repository

import (
	"context"
	"database/sql"
	"golang-simple-api/entity"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer
	GetAllCustomer(ctx context.Context, tx *sql.Tx) []entity.Customer
	GetCustomerById(ctx context.Context, tx *sql.Tx, customerId int) (entity.Customer, error)
	UpdateCustomer(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer
	DeleteCustomer(ctx context.Context, tx *sql.Tx, customerId int)
	GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.Customer, error)
}
