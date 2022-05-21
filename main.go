package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-simple-api/config"
	"golang-simple-api/controller"
	"golang-simple-api/exception"
	"golang-simple-api/repository"
	"golang-simple-api/service"
)

func main() {
	configuration := config.New(".env")
	database := config.NewMySQLDatabase(configuration)

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(database, &customerRepository)
	customerController := controller.NewCustomerController(&customerService)

	app := echo.New()
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{DisablePrintStack: true}))
	app.HTTPErrorHandler = exception.ErrorHandler
	customerController.Route(app)
	err := app.Start(":9000")
	exception.PanicIfNeeded(err)
}
