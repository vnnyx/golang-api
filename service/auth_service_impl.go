package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vnnyx/golang-api/config"
	"github.com/vnnyx/golang-api/exception"
	"github.com/vnnyx/golang-api/helper"
	"github.com/vnnyx/golang-api/model"
	"github.com/vnnyx/golang-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repository.CustomerRepository
	repository.AuthRepository
	*config.Config
	*sqlx.DB
}

func NewAuthService(config *config.Config, DB *sqlx.DB) AuthService {
	return &AuthServiceImpl{Config: config, DB: DB}
}

func (service *AuthServiceImpl) InjectCustomerRepository(customerRepository repository.CustomerRepository) {
	service.CustomerRepository = customerRepository
}

func (service *AuthServiceImpl) InjectAuthRepository(authRepository repository.AuthRepository) {
	service.AuthRepository = authRepository
}

func (service *AuthServiceImpl) Login(ctx context.Context, request model.LoginRequest) (response model.LoginResponse, err error) {
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return model.LoginResponse{}, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	customer, err := service.CustomerRepository.GetUserByUsername(ctx, tx, request.Username)
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
	}, service.Config)

	tokenDetails := &model.TokenDetails{
		AccessToken: td.AccessToken,
		AccessUuid:  td.AccessUuid,
		AtExpires:   td.AtExpires,
	}

	err = service.AuthRepository.StoreToken(ctx, *tokenDetails)
	exception.PanicIfNeeded(err)

	response = model.LoginResponse{
		AccessToken: td.AccessToken,
		UserId:      customer.Id,
		Username:    customer.Username,
		Email:       customer.Email,
	}

	return response, nil
}

func (service *AuthServiceImpl) Logout(ctx context.Context, accessUuid string) {
	err := service.AuthRepository.DeleteToken(ctx, accessUuid)
	if err != nil {
		fmt.Println(err)
		exception.PanicIfNeeded(errors.New(model.UNAUTHORIZATION))
	}

}
