package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vnnyx/golang-api/config"
	"github.com/vnnyx/golang-api/controller"
	"github.com/vnnyx/golang-api/exception"
	"github.com/vnnyx/golang-api/repository"
	"github.com/vnnyx/golang-api/service"
)

func main() {
	configuration, err := config.NewConfig(".", ".env")
	rdb := config.NewRedisClient()
	database := config.NewMySQLDatabase(configuration)

	customerRepository := repository.NewCustomerRepository()
	authRepository := repository.NewAuthRepository(rdb)
	customerService := service.NewCustomerService(customerRepository, database)
	authService := service.NewAuthService(configuration, database)
	authService.InjectAuthRepository(authRepository)
	authService.InjectCustomerRepository(customerRepository)
	customerController := controller.NewCustomerController(&customerService)
	authController := controller.NewAuthController(&authService)

	app := echo.New()
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{DisablePrintStack: true}))
	app.Use(middleware.CORS())
	app.HTTPErrorHandler = exception.ErrorHandler
	customerController.Route(app)
	authController.Route(app)
	err = app.Start(fmt.Sprintf(":%v", configuration.AppPort))
	exception.PanicIfNeeded(err)
}
