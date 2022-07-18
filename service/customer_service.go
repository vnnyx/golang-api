package service

import (
	"context"
	"github.com/vnnyx/golang-api/model"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, request model.CustomerCreateRequest) (model.CustomerResponse, error)
	GetAllCustomer(ctx context.Context) ([]model.CustomerResponse, error)
	GetCustomerById(ctx context.Context, customerId uint32) (model.CustomerResponse, error)
	UpdateCustomer(ctx context.Context, request model.CustomerUpdateRequest) (model.CustomerResponse, error)
	DeleteCustomer(ctx context.Context, customerId uint32) error
}
