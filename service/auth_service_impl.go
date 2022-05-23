package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang-simple-api/exception"
	"golang-simple-api/helper"
	"golang-simple-api/model"
	"golang-simple-api/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthServiceImpl struct {
	DB *sql.DB
	repository.CustomerRepository
	Redis *redis.Client
}

func NewAuthService(DB *sql.DB, customerRepository *repository.CustomerRepository, Redis *redis.Client) AuthService {
	return &AuthServiceImpl{DB: DB, CustomerRepository: *customerRepository, Redis: Redis}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request model.LoginRequest) (response model.LoginResponse, err error) {
	fmt.Println("AUTH_SERVICE")
	tx, err := service.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.GetUserByUsername(ctx, tx, request.Username)
	if err != nil {
		return response, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New(model.UNAUTHORIZATION)
	}

	td := helper.CreateToken(model.JwtPayload{
		UserId:   customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
	})

	service.Redis.Set(ctx, td.AccessUuid, td.AccessToken, time.Unix(td.AtExpires, 0).Sub(time.Now()))

	response = model.LoginResponse{
		AccessToken: td.AccessToken,
		UserId:      customer.Id,
		Username:    customer.Username,
		Email:       customer.Email,
	}

	return response, nil
}

func (service *AuthServiceImpl) Logout(ctx context.Context, accessUuid string) {
	err := service.Redis.Del(ctx, accessUuid).Err()
	if err != nil {
		exception.PanicIfNeeded(model.UNAUTHORIZATION)
	}

}
