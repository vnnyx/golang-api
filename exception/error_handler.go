package exception

import (
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/vnnyx/golang-api/model"
	"net/http"
	"strings"
)

func ErrorHandler(err error, ctx echo.Context) {

	if databaseError(err, ctx) {
		return
	}
	if validationError(err, ctx) {
		return
	}
	generalError(err, ctx)
}

func generalError(err error, ctx echo.Context) {
	switch err.Error() {
	case "USER_NOT_FOUND":
		_ = ctx.JSON(http.StatusNotFound, model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"id": "NOT_FOUND",
			},
		})
		break
	case model.UNAUTHORIZATION:
		_ = ctx.JSON(http.StatusUnauthorized, model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   nil,
			Error: map[string]interface{}{
				"general": "UNAUTHORIZED",
			},
		})
		break
	default:
		_ = ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   nil,
			Error: map[string]interface{}{
				"message": "Internal server error",
			},
		})
	}
}

func validationError(err error, ctx echo.Context) bool {
	_, ok := err.(ValidationError)
	if ok {
		_ = ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error:  err.Error(),
		})
		return true
	}
	return false
}

func databaseError(err error, ctx echo.Context) bool {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}
	if sqlError.Number == 1062 && strings.Contains(sqlError.Message, "username") {
		_ = ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"username": "MUST_UNIQUE",
			},
		})
	} else if sqlError.Number == 1062 && strings.Contains(sqlError.Message, "email") {
		_ = ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"email": "MUST_UNIQUE",
			},
		})
	} else {
		_ = ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   nil,
			Error: map[string]interface{}{
				"message": "Internal server error",
			},
		})
	}
	return true
}
