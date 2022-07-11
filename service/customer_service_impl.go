package service

import (
	"context"
	"golang-simple-api/entity"
	"golang-simple-api/exception"
	"golang-simple-api/model"
	"golang-simple-api/repository"
	"golang-simple-api/validation"
	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceImpl struct {
	repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{CustomerRepository: customerRepository}
}

func (service *CustomerServiceImpl) CreateCustomer(ctx context.Context, request model.CustomerCreateRequest) model.CustomerResponse {
	validation.Validate(request)

	customer := entity.Customer{
		Id:       request.Id,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
	}

	customer = service.CustomerRepository.CreateCustomer(ctx, customer)

	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}

	return response
}

func (service *CustomerServiceImpl) GetAllCustomer(ctx context.Context) []model.CustomerResponse {
	customers := service.CustomerRepository.GetAllCustomer(ctx)
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

func (service *CustomerServiceImpl) GetCustomerById(ctx context.Context, customerId uint32) model.CustomerResponse {
	customer, err := service.CustomerRepository.GetCustomerById(ctx, customerId)
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

	customer, err := service.CustomerRepository.GetCustomerById(ctx, request.Id)
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

	customer = service.CustomerRepository.UpdateCustomer(ctx, customer)

	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}

	return response
}

func (service *CustomerServiceImpl) DeleteCustomer(ctx context.Context, customerId uint32) {
	_, err := service.CustomerRepository.GetCustomerById(ctx, customerId)
	exception.PanicIfNeeded(err)

	service.CustomerRepository.DeleteCustomer(ctx, customerId)
}
