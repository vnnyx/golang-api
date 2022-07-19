package controller

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/vnnyx/golang-api/entity"
	"github.com/vnnyx/golang-api/model"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCustomerController_CreateCustomer_Success(t *testing.T) {
	tx, err := database.BeginTxx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = customerRepository.DeleteAllCustomer(context.Background(), tx)

	customerCreateRequest := model.CustomerCreateRequest{
		Username: "username_test",
		Email:    "test@email.com",
		Password: "password",
		Gender:   "test",
	}
	requestBody, _ := json.Marshal(customerCreateRequest)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	request := httptest.NewRequest("POST", "/api/customer", bytes.NewBuffer(requestBody))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()

	app.ServeHTTP(recorder, request)
	response := recorder.Result()

	//Check response body
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	customerResponse := model.CustomerResponse{}
	json.Unmarshal(jsonData, &customerResponse)
	assert.NotNil(t, customerResponse)
	assert.Equal(t, customerCreateRequest.Username, customerResponse.Username)
}

func TestCustomerController_GetAllCustomer_Success(t *testing.T) {
	tx, err := database.BeginTxx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = customerRepository.DeleteAllCustomer(context.Background(), tx)
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	customerEntity := entity.Customer{
		Id:       uint32(1),
		Username: "username_test",
		Email:    "test@email.com",
		Password: string(password),
		Gender:   "test",
	}

	customerRepository.CreateCustomer(context.Background(), tx, customerEntity)

	authRepository.StoreToken(context.Background(), *tokenDetails)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	request := httptest.NewRequest("GET", "/api/customer", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	request.Header.Set("Authorization", "Bearer "+tokenDetails.AccessToken)

	recorder := httptest.NewRecorder()

	app.ServeHTTP(recorder, request)

	response := recorder.Result()

	//Check response body
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	customerResponse := webResponse.Data.([]interface{})
	getCustomerResponse := model.CustomerResponse{}
	containsCustomer := false
	for _, data := range customerResponse {
		jsonData, _ := json.Marshal(data)
		json.Unmarshal(jsonData, &getCustomerResponse)
		if getCustomerResponse.Id == customerEntity.Id {
			containsCustomer = true
		}
	}
	assert.True(t, containsCustomer)
}

func TestCustomerController_UpdateCustomer_Success(t *testing.T) {
	tx, err := database.BeginTxx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = customerRepository.DeleteAllCustomer(context.Background(), tx)
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	customerEntity := entity.Customer{
		Id:       uint32(1),
		Username: "username_test",
		Email:    "test@email.com",
		Password: string(password),
		Gender:   "test",
	}

	customerRepository.CreateCustomer(context.Background(), tx, customerEntity)
	authRepository.StoreToken(context.Background(), *tokenDetails)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	customerUpdateRequest := model.CustomerUpdateRequest{
		Id:       customerEntity.Id,
		Username: "username_update",
		Email:    "email_update",
		Password: string(password),
		Gender:   customerEntity.Gender,
	}

	requestBody, _ := json.Marshal(customerUpdateRequest)

	request := httptest.NewRequest("PUT", "/api/customer/"+strconv.Itoa(int(customerEntity.Id)), bytes.NewBuffer(requestBody))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	request.Header.Set("Authorization", "Bearer "+tokenDetails.AccessToken)

	recorder := httptest.NewRecorder()

	app.ServeHTTP(recorder, request)

	response := recorder.Result()

	//Check http status
	assert.Equal(t, 200, response.StatusCode)

	//Check response body
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	customerResponse := model.CustomerResponse{}
	json.Unmarshal(jsonData, &customerResponse)
	assert.NotNil(t, customerResponse)
	assert.Equal(t, customerEntity.Id, customerResponse.Id)
	assert.NotEqual(t, customerEntity.Username, customerResponse.Username)
	assert.NotEqual(t, customerEntity.Email, customerResponse.Email)
}

func TestCustomerController_GetCustomerById_Success(t *testing.T) {
	tx, err := database.BeginTxx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = customerRepository.DeleteAllCustomer(context.Background(), tx)
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	customerEntity := entity.Customer{
		Id:       uint32(1),
		Username: "username_test",
		Email:    "test@email.com",
		Password: string(password),
		Gender:   "test",
	}

	customerRepository.CreateCustomer(context.Background(), tx, customerEntity)
	authRepository.StoreToken(context.Background(), *tokenDetails)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	request := httptest.NewRequest("GET", "/api/customer/"+strconv.Itoa(int(customerEntity.Id)), nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	request.Header.Set("Authorization", "Bearer "+tokenDetails.AccessToken)

	recorder := httptest.NewRecorder()

	app.ServeHTTP(recorder, request)

	response := recorder.Result()

	//Check http status
	assert.Equal(t, 200, response.StatusCode)

	//Check response body
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	customerResponse := model.CustomerResponse{}
	json.Unmarshal(jsonData, &customerResponse)
	assert.NotNil(t, customerResponse)
	assert.Equal(t, customerEntity.Id, customerResponse.Id)
	assert.Equal(t, customerEntity.Username, customerResponse.Username)
}

func TestCustomerController_DeleteCustomer_Success(t *testing.T) {
	tx, err := database.BeginTxx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = customerRepository.DeleteAllCustomer(context.Background(), tx)
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	customerEntity := entity.Customer{
		Id:       uint32(1),
		Username: "username_test",
		Email:    "test@email.com",
		Password: string(password),
		Gender:   "test",
	}

	customerRepository.CreateCustomer(context.Background(), tx, customerEntity)
	authRepository.StoreToken(context.Background(), *tokenDetails)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	request := httptest.NewRequest("DELETE", "/api/customer/"+strconv.Itoa(int(customerEntity.Id)), nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	request.Header.Set("Authorization", "Bearer "+tokenDetails.AccessToken)

	recorder := httptest.NewRecorder()

	app.ServeHTTP(recorder, request)

	response := recorder.Result()

	//Check http status
	assert.Equal(t, 200, response.StatusCode)

	//Check response body
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.Nil(t, webResponse.Data)
}
