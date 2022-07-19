package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/vnnyx/golang-api/entity"
	"github.com/vnnyx/golang-api/model"
	"github.com/vnnyx/golang-api/repository"
	"github.com/vnnyx/golang-api/validation"
)

type CustomerServiceImpl struct {
	repository.CustomerRepository
	DB *sqlx.DB
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sqlx.DB) CustomerService {
	return &CustomerServiceImpl{CustomerRepository: customerRepository, DB: DB}
}

func (service *CustomerServiceImpl) CreateCustomer(ctx context.Context, request model.CustomerCreateRequest) (model.CustomerResponse, error) {
	validation.Validate(request)
	customer := entity.Customer{
		Id:       request.Id,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
	}

	tx, err := service.DB.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})

	if err != nil {
		return model.CustomerResponse{}, err
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
	customer, err = service.CustomerRepository.CreateCustomer(ctx, tx, customer)
	if err != nil {
		return model.CustomerResponse{}, err
	}
	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}
	return response, nil
}

func (service *CustomerServiceImpl) GetAllCustomer(ctx context.Context) ([]model.CustomerResponse, error) {
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return []model.CustomerResponse{}, err
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
	customers, err := service.CustomerRepository.GetAllCustomer(ctx, tx)
	if err != nil {
		return []model.CustomerResponse{}, err
	}
	var response []model.CustomerResponse
	for _, customer := range customers {
		response = append(response, model.CustomerResponse{
			Id:       customer.Id,
			Username: customer.Username,
			Email:    customer.Email,
			Gender:   customer.Gender,
		})
	}
	return response, nil
}

func (service *CustomerServiceImpl) GetCustomerById(ctx context.Context, customerId uint32) (model.CustomerResponse, error) {
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return model.CustomerResponse{}, err
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
	customer, err := service.CustomerRepository.GetCustomerById(ctx, tx, customerId)
	if err != nil {
		return model.CustomerResponse{}, errors.New("USER_NOT_FOUND")
	}

	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}
	return response, nil
}

func (service *CustomerServiceImpl) UpdateCustomer(ctx context.Context, request model.CustomerUpdateRequest) (model.CustomerResponse, error) {
	validation.UpdateValidate(request)

	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return model.CustomerResponse{}, err
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
	customer, err := service.CustomerRepository.GetCustomerById(ctx, tx, request.Id)
	if err != nil {
		return model.CustomerResponse{}, errors.New("USER_NOT_FOUND")
	}

	customer = entity.Customer{
		Id:       request.Id,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
	}

	customer, err = service.CustomerRepository.UpdateCustomer(ctx, tx, customer)
	if err != nil {
		return model.CustomerResponse{}, err
	}

	response := model.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Email:    customer.Email,
		Gender:   customer.Gender,
	}

	return response, nil
}

func (service *CustomerServiceImpl) DeleteCustomer(ctx context.Context, customerId uint32) error {
	tx, err := service.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
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
	_, err = service.CustomerRepository.GetCustomerById(ctx, tx, customerId)
	if err != nil {
		return err
	}
	err = service.CustomerRepository.DeleteCustomer(ctx, tx, customerId)
	if err != nil {
		return err
	}
	return nil
}
