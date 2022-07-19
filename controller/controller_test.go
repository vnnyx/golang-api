package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vnnyx/golang-api/config"
	"github.com/vnnyx/golang-api/exception"
	"github.com/vnnyx/golang-api/helper"
	"github.com/vnnyx/golang-api/model"
	"github.com/vnnyx/golang-api/repository"
	"github.com/vnnyx/golang-api/service"
)

var (
	configuration, _ = config.NewConfig("../", ".env.test")
	rdb              = config.NewRedisClient()
	database         = config.NewMySQLDatabase(configuration)

	customerRepository = repository.NewCustomerRepository()
	authRepository     = repository.NewAuthRepository(rdb)
	customerService    = service.NewCustomerService(customerRepository, database)
	authService        = service.NewAuthService(configuration, database)
	customerController = NewCustomerController(&customerService)
	authController     = NewAuthController(&authService)
	app                = createTestApp()
	td                 = helper.CreateToken(model.JwtPayload{
		UserId:   1,
		Username: "username_test",
		Email:    "test@email.com",
	}, configuration)

	tokenDetails = &model.TokenDetails{
		AccessToken: td.AccessToken,
		AccessUuid:  td.AccessUuid,
		AtExpires:   td.AtExpires,
	}
)

func createTestApp() *echo.Echo {
	var app = echo.New()
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{DisablePrintStack: true}))
	app.Use(middleware.CORS())
	app.HTTPErrorHandler = exception.ErrorHandler
	customerController.Route(app)
	authController.Route(app)
	return app
}
