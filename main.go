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
	rdb := config.NewRedisClient()
	database := config.NewMySQLDatabase(configuration)

	customerRepository := repository.NewCustomerRepository(database)
	customerService := service.NewCustomerService(&customerRepository)
	authService := service.NewAuthService(&customerRepository, rdb)
	customerController := controller.NewCustomerController(&customerService)
	authController := controller.NewAuthController(&authService)

	app := echo.New()
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{DisablePrintStack: true}))
	app.Use(middleware.CORS())
	app.HTTPErrorHandler = exception.ErrorHandler
	customerController.Route(app)
	authController.Route(app)
	err := app.Start(":9000")
	exception.PanicIfNeeded(err)
}
