package service

import (
	"context"
	"golang-simple-api/model"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, request model.CustomerCreateRequest) model.CustomerResponse
	GetAllCustomer(ctx context.Context) []model.CustomerResponse
	GetCustomerById(ctx context.Context, customerId uint32) model.CustomerResponse
	UpdateCustomer(ctx context.Context, request model.CustomerUpdateRequest) model.CustomerResponse
	DeleteCustomer(ctx context.Context, customerId uint32)
}
