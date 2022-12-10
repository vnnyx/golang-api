package service

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vnnyx/golang-api/entity"
	"github.com/vnnyx/golang-api/model"
	"testing"
)

func TestCustomerServiceImpl_CreateCustomer(t *testing.T) {
	defer db.Close()
	tests := []struct {
		name     string
		repoName string
		error    error
	}{
		{
			name:     "TestCustomerServiceImpl_CreateCustomer_Success",
			repoName: "CreateCustomer",
			error:    nil,
		},
		{
			name:     "TestCustomerServiceImpl_CreateCustomer_Failed",
			repoName: "CreateCustomer",
			error:    errors.New("error_test"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockx.ExpectBegin()
			customerRepoMock.On(test.repoName, ctx, mock.Anything, customerEntity).Return(customerEntity, test.error)
			if test.error != nil {
				mockx.ExpectRollback()
			} else {
				mockx.ExpectCommit()
			}

			response, err := customerService.CreateCustomer(ctx, model.CustomerCreateRequest{
				Id:       customerEntity.Id,
				Username: customerEntity.Username,
				Email:    customerEntity.Email,
				Password: customerEntity.Password,
				Gender:   customerEntity.Gender,
			})

			if err != nil {
				assert.NotNil(t, err)
				assert.NotEqual(t, customerEntity.Id, response.Id)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

func TestCustomerServiceImpl_GetAllCustomer(t *testing.T) {
	defer db.Close()
	tests := []struct {
		name       string
		repoName   string
		repoResult []entity.Customer
		error      error
	}{
		{
			name:       "TestCustomerServiceImpl_GetAllCustomer_Success",
			repoName:   "GetAllCustomer",
			repoResult: []entity.Customer{customerEntity},
			error:      nil,
		},
		{
			name:       "TestCustomerServiceImpl_GetAllCustomer_Failed",
			repoName:   "GetAllCustomer",
			repoResult: []entity.Customer{},
			error:      errors.New("error_test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockx.ExpectBegin()
			customerRepoMock.On(test.repoName, ctx, mock.Anything).Return(test.repoResult, test.error)
			if test.error != nil {
				mockx.ExpectRollback()
			} else {
				mockx.ExpectCommit()
			}

			response, err := customerService.GetAllCustomer(ctx)
			if err != nil {
				assert.NotNil(t, err)
				assert.NotNil(t, response)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

func TestCustomerServiceImpl_GetCustomerById(t *testing.T) {
	defer db.Close()
	tests := []struct {
		name       string
		repoName   string
		customerId uint32
		error      error
	}{
		{
			name:       "TestCustomerServiceImpl_GetCustomerById_Success",
			repoName:   "GetCustomerById",
			customerId: customerEntity.Id,
			error:      nil,
		},
		{
			name:       "TestCustomerServiceImpl_GetCustomerById_Failed",
			repoName:   "GetCustomerById",
			customerId: uint32(2),
			error:      errors.New("error_test"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockx.ExpectBegin()
			customerRepoMock.On(test.repoName, ctx, mock.Anything, test.customerId).Return(customerEntity, test.error)
			if test.error != nil {
				mockx.ExpectRollback()
			} else {
				mockx.ExpectCommit()
			}

			response, err := customerService.GetCustomerById(ctx, test.customerId)
			if err != nil {
				assert.NotNil(t, err)
				assert.NotEqual(t, customerEntity.Id, response.Id)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

func TestCustomerServiceImpl_UpdateCustomer(t *testing.T) {
	defer db.Close()
	tests := []struct {
		name       string
		repoName   string
		customerId uint32
		error      error
	}{
		{
			name:       "TestCustomerServiceImpl_UpdateCustomer_Success",
			repoName:   "UpdateCustomer",
			customerId: customerEntity.Id,
			error:      nil,
		},
		{
			name:       "TestCustomerServiceImpl_UpdateCustomer_Failed",
			repoName:   "UpdateCustomer",
			customerId: uint32(1),
			error:      errors.New("error_test"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockx.ExpectBegin()
			customerRepoMock.On("GetCustomerById", ctx, mock.Anything, test.customerId).Return(customerEntity, test.error)
			customerRepoMock.On(test.repoName, ctx, mock.Anything, customerEntity).Return(customerEntity, test.error)
			if test.error != nil {
				mockx.ExpectRollback()
			} else {
				mockx.ExpectCommit()
			}
			response, err := customerService.UpdateCustomer(ctx, model.CustomerUpdateRequest{
				Id:       customerEntity.Id,
				Username: customerEntity.Username,
				Email:    customerEntity.Email,
				Password: customerEntity.Password,
				Gender:   customerEntity.Gender,
			})
			if err != nil {
				assert.NotNil(t, err)
				assert.NotEqual(t, customerEntity.Id, response.Id)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

func TestCustomerServiceImpl_DeleteCustomer(t *testing.T) {
	defer db.Close()
	tests := []struct {
		name       string
		repoName   string
		error      error
		customerId uint32
	}{
		{
			name:       "TestCustomerServiceImpl_DeleteCustomer_Success",
			repoName:   "DeleteCustomer",
			error:      nil,
			customerId: customerEntity.Id,
		},
		{
			name:       "TestCustomerServiceImpl_DeleteCustomer_Failed",
			repoName:   "DeleteCustomer",
			error:      errors.New("error_test"),
			customerId: uint32(2),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockx.ExpectBegin()
			customerRepoMock.On("GetCustomerById", ctx, mock.Anything, test.customerId).Return(customerEntity, test.error)
			customerRepoMock.On(test.repoName, ctx, mock.Anything, test.customerId).Return(test.error)
			if test.error != nil {
				mockx.ExpectRollback()
			} else {
				mockx.ExpectCommit()
			}

			err := customerService.DeleteCustomer(ctx, test.customerId)
			if err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
