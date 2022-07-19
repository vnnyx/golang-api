package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/vnnyx/golang-api/middleware"
	"github.com/vnnyx/golang-api/model"
	"github.com/vnnyx/golang-api/service"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

func (controller AuthController) Route(e *echo.Echo) {
	router := e.Group("/api/auth")
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout, middleware.CheckToken)
}

func (controller AuthController) Login(c echo.Context) error {
	var request model.LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	response, err := controller.AuthService.Login(c.Request().Context(), request)
	if err != nil {
		return err
	}

	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller AuthController) Logout(c echo.Context) error {
	accessUuid := c.Get("currentAccessUuid")
	err := controller.AuthService.Logout(c.Request().Context(), accessUuid.(string))
	if err != nil {
		return err
	}
	return c.JSON(200, model.WebResponse{
		Code:   200,
		Status: "OK",
	})
}
