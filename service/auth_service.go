package service

import (
	"context"
	"golang-simple-api/model"
)

type AuthService interface {
	Login(ctx context.Context, request model.LoginRequest) (response model.LoginResponse, err error)
	Logout(ctx context.Context, accessUuid string)
}
