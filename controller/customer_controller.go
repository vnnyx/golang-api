package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang-simple-api/exception"
	"golang-simple-api/middleware"
	"golang-simple-api/model"
	"golang-simple-api/service"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type CustomerController struct {
	service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) CustomerController {
	return CustomerController{CustomerService: *customerService}
}

func (controller *CustomerController) Route(e *echo.Echo) {
	e.POST("/api/customer", controller.CreateCustomer)
	e.GET("/api/customer/:id", controller.GetCustomerById, middleware.CheckToken)
	e.GET("/api/customer", controller.GetAllCustomer, middleware.CheckToken)
	e.PUT("/api/customer/:id", controller.UpdateCustomer, middleware.CheckToken)
	e.DELETE("/api/customer/:id", controller.DeleteCustomer, middleware.CheckToken)
}

func (controller CustomerController) CreateCustomer(c echo.Context) error {
	var request model.CustomerCreateRequest
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	exception.PanicIfNeeded(err)
	request.Password = string(password)
	request.Id = uuid.New().ID()

	response := controller.CustomerService.CreateCustomer(c.Request().Context(), request)
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller CustomerController) GetCustomerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response := controller.CustomerService.GetCustomerById(c.Request().Context(), uint32(id))
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller CustomerController) GetAllCustomer(c echo.Context) error {
	response := controller.CustomerService.GetAllCustomer(c.Request().Context())
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller CustomerController) UpdateCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var request model.CustomerUpdateRequest
	request.Id = uint32(id)
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)
	response := controller.CustomerService.UpdateCustomer(c.Request().Context(), request)
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller CustomerController) DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	controller.CustomerService.DeleteCustomer(c.Request().Context(), uint32(id))
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
	})
}
