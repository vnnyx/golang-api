package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAuthServiceImpl_Login(t *testing.T) {
	defer db.Close()
	tests := []struct {
		name  string
		error error
	}{
		{
			name:  "TestAuthServiceImpl_Login_Success",
			error: nil,
		},
		{
			name:  "TestAuthServiceImpl_Login_Failed",
			error: errors.New("error_test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockx.ExpectBegin()
			customerRepoMock.On("GetUserByUsername", ctx, mock.Anything, loginRequest.Username).Return(customerEntity, test.error)
			if test.error != nil {
				mockx.ExpectRollback()
			} else {
				mockx.ExpectCommit()
			}
			authRepoMock.On("StoreToken", ctx, mock.Anything).Return(test.error)

			authService.InjectAuthRepository(authRepoMock)
			authService.InjectCustomerRepository(customerRepoMock)
			response, err := authService.Login(ctx, loginRequest)
			if err != nil {
				assert.NotNil(t, err)
				assert.NotEqual(t, customerEntity.Username, response.Username)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, customerEntity.Username, response.Username)
			}

		})
	}
}

func TestAuthServiceImpl_Logout(t *testing.T) {
	tests := []struct {
		name  string
		error error
	}{
		{
			name:  "TestAuthServiceImpl_Logout_Success",
			error: nil,
		},
		{
			name:  "TestAuthServiceImpl_Logout_Failed",
			error: errors.New("error_test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authRepoMock.On("DeleteToken", mock.Anything).Return(test.error)
			authService.InjectAuthRepository(authRepoMock)
			err := authService.Logout(ctx, mock.Anything)
			if err != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
