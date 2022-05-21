package exception

import (
	"github.com/labstack/echo/v4"
	"golang-simple-api/model"
	"net/http"
)

func ErrorHandler(err error, ctx echo.Context) {
	if err.Error() == "USERNAME_REGISTERED" {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"username": "MUST_UNIQUE",
			},
		})
	} else if err.Error() == "USER_NOT_FOUND" {
		ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"id": "NOT_FOUND",
			},
		})

	} else {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

}
