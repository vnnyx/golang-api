package service

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/vnnyx/golang-api/config"
	"github.com/vnnyx/golang-api/entity"
	"github.com/vnnyx/golang-api/model"
	"github.com/vnnyx/golang-api/repository/mocks"
	"golang.org/x/crypto/bcrypt"
)

var (
	configuration, _ = config.NewConfig("../", ".env")
	hashPassword, _  = bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	customerEntity   = entity.Customer{
		Id:       1,
		Username: "username_test",
		Email:    "email_test",
		Password: string(hashPassword),
		Gender:   "test",
	}
	loginRequest = model.LoginRequest{
		Username: "username_test",
		Password: "password",
	}
	//Context
	ctx = context.Background()
	//Setup DB
	db, mockx, _    = sqlmock.New()
	dbx             = sqlx.NewDb(db, "sqlmock")
	customerService = NewCustomerService(customerRepoMock, dbx)
	authService     = NewAuthService(configuration, dbx)
	//Setup repo
	authRepoMock     = new(mocks.AuthRepository)
	customerRepoMock = new(mocks.CustomerRepository)
)
