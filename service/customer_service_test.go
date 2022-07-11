package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang-simple-api/entity"
	"golang-simple-api/model"
	"golang-simple-api/repository/mocks"
	"testing"
)

var (
	customerEntity = entity.Customer{
		Id:       1,
		Username: "username_test",
		Email:    "email_test",
		Password: "password",
		Gender:   "test",
	}
	customerRepoMock = new(mocks.CustomerRepository)
	customerService  = NewCustomerService(customerRepoMock)
)

func TestCustomerServiceImpl_CreateCustomer(t *testing.T) {
	customerRepoMock.On("CreateCustomer", context.Background(), customerEntity).Return(customerEntity)

	response := customerService.CreateCustomer(context.Background(), model.CustomerCreateRequest{
		Id:       customerEntity.Id,
		Username: customerEntity.Username,
		Email:    customerEntity.Email,
		Password: customerEntity.Password,
		Gender:   customerEntity.Gender,
	})
	assert.Equal(t, customerEntity.Username, response.Username)
}

func TestCustomerServiceImpl_GetAllCustomer(t *testing.T) {
	customerRepoMock.On("GetAllCustomer", context.Background()).Return([]entity.Customer{customerEntity})

	response := customerService.GetAllCustomer(context.Background())
	assert.Equal(t, customerEntity.Username, response[0].Username)
}

func TestCustomerServiceImpl_GetCustomerById(t *testing.T) {
	customerRepoMock.On("GetCustomerById", context.Background(), customerEntity.Id).Return(customerEntity, nil)

	response := customerService.GetCustomerById(context.Background(), customerEntity.Id)
	assert.Equal(t, customerEntity.Username, response.Username)
}
