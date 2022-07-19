package service

import (
	"context"
	"github.com/vnnyx/golang-api/model"
	"github.com/vnnyx/golang-api/repository"
)

type AuthService interface {
	InjectCustomerRepository(customerRepository repository.CustomerRepository)
	InjectAuthRepository(authRepository repository.AuthRepository)
	Login(ctx context.Context, request model.LoginRequest) (response model.LoginResponse, err error)
	Logout(ctx context.Context, accessUuid string) error
}
