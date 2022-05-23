package controller

import (
	"github.com/labstack/echo/v4"
	"golang-simple-api/exception"
	"golang-simple-api/middleware"
	"golang-simple-api/model"
	"golang-simple-api/service"
	"strconv"
)

type CustomerController struct {
	service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) CustomerController {
	return CustomerController{CustomerService: *customerService}
}

func (controller *CustomerController) Route(e *echo.Echo) {
	router := e.Group("/api/customer", middleware.CheckToken)
	router.POST("", controller.CreateCustomer)
	router.GET("/:id", controller.GetCustomerById)
	router.GET("", controller.GetAllCustomer)
	router.PUT("/:id", controller.UpdateCustomer)
	router.DELETE("/:id", controller.DeleteCustomer)
}

func (controller CustomerController) CreateCustomer(c echo.Context) error {
	var request model.CustomerCreateRequest
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)

	response := controller.CustomerService.CreateCustomer(c.Request().Context(), request)
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller CustomerController) GetCustomerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response := controller.CustomerService.GetCustomerById(c.Request().Context(), id)
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
	request.Id = id
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
	controller.CustomerService.DeleteCustomer(c.Request().Context(), id)
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
	})
}
