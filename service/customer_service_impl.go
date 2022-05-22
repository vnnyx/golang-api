package service

import (
	"context"
	"database/sql"
	"golang-simple-api/entity"
	"golang-simple-api/exception"
	"golang-simple-api/helper"
	"golang-simple-api/model"
	"golang-simple-api/repository"
	"golang-simple-api/validation"
	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceImpl struct {
	DB *sql.DB
	repository.CustomerRepository
}

func NewCustomerService(DB *sql.DB, customerRepository *repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{DB: DB, CustomerRepository: *customerRepository}
}

func (service *CustomerServiceImpl) CreateCustomer(ctx context.Context, request model.CustomerCreateRequest) model.CustomerResponse {
	validation.Validate(request)

	tx, err := service.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	exception.PanicIfNeeded(err)

	customer := entity.Customer{
		Username: request.Username,
		Email:    request.Email,
		Password: string(password),
		Gender:   request.Gender,
	}

	customer = service.CustomerRepository.CreateCustomer(ctx, tx, customer)

	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}

	return response
}

func (service *CustomerServiceImpl) GetAllCustomer(ctx context.Context) []model.CustomerResponse {
	tx, err := service.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.GetAllCustomer(ctx, tx)
	var response []model.CustomerResponse
	for _, customer := range customers {
		response = append(response, model.CustomerResponse{
			Id:       customer.Id,
			Username: customer.Username,
			Email:    customer.Email,
			Gender:   customer.Gender,
		})
	}
	return response
}

func (service *CustomerServiceImpl) GetCustomerById(ctx context.Context, customerId int) model.CustomerResponse {
	tx, err := service.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.GetCustomerById(ctx, tx, customerId)
	if err != nil {
		exception.PanicIfNeeded("USER_NOT_FOUND")
	}

	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}
	return response
}

func (service *CustomerServiceImpl) UpdateCustomer(ctx context.Context, request model.CustomerUpdateRequest) model.CustomerResponse {
	validation.UpdateValidate(request)
	tx, err := service.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.GetCustomerById(ctx, tx, request.Id)
	if err != nil {
		exception.PanicIfNeeded("USER_NOT_FOUND")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	exception.PanicIfNeeded(err)
	customer = entity.Customer{
		Id:       request.Id,
		Username: request.Username,
		Email:    request.Email,
		Password: string(password),
		Gender:   request.Gender,
	}

	customer = service.CustomerRepository.UpdateCustomer(ctx, tx, customer)

	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}

	return response
}

func (service *CustomerServiceImpl) DeleteCustomer(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.CustomerRepository.GetCustomerById(ctx, tx, customerId)
	exception.PanicIfNeeded(err)

	service.CustomerRepository.DeleteCustomer(ctx, tx, customerId)
}
