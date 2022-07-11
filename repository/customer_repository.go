package repository

import (
	"context"
	"golang-simple-api/entity"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer entity.Customer) entity.Customer
	GetAllCustomer(ctx context.Context) []entity.Customer
	GetCustomerById(ctx context.Context, customerId int) (entity.Customer, error)
	UpdateCustomer(ctx context.Context, customer entity.Customer) entity.Customer
	DeleteCustomer(ctx context.Context, customerId int)
	GetUserByUsername(ctx context.Context, username string) (entity.Customer, error)
}
