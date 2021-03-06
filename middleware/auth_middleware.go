package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/vnnyx/golang-api/config"
	"github.com/vnnyx/golang-api/exception"
	"github.com/vnnyx/golang-api/model"
	"github.com/vnnyx/golang-api/repository"
	"net/http"
	"strings"
)

type DecodedStructure struct {
	UserId     int    `json:"user_id"`
	Username   string `json:"username"`
	AccessUuid string `json:"access_uuid"`
}

func ValidateToken(encodedToken string) (token *jwt.Token, errData error) {
	configuration, err := config.NewConfig(".", ".env")
	exception.PanicIfNeeded(err)
	jwtPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(configuration.JWTPublicKey))

	if err != nil {
		return token, err
	}

	tokenString := encodedToken
	claims := jwt.MapClaims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPublicKey, nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, errors.New("invalid token")
	}
	return token, nil
}

func DecodeToken(encodedToken string) (decodedResult DecodedStructure, errData error) {
	configuration, err := config.NewConfig(".", ".env")
	exception.PanicIfNeeded(err)
	jwtPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(configuration.JWTPublicKey))
	tokenString := encodedToken
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPublicKey, nil
	})
	if err != nil {
		return decodedResult, err
	}
	if !token.Valid {
		return decodedResult, errors.New("invalid token")
	}

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return decodedResult, err
	}

	var obj DecodedStructure
	if err := json.Unmarshal(jsonbody, &obj); err != nil {
		return decodedResult, err
	}

	return obj, nil
}

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		header := ctx.Request().Header
		tokenSlice := strings.Split(header.Get("Authorization"), "Bearer ")

		var tokenString string
		if len(tokenSlice) == 2 {
			tokenString = tokenSlice[1]
		}

		//validate token
		_, err := ValidateToken(tokenString)
		if err != nil {
			fmt.Println("validate")
			return ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			})
		}

		//extract data from token
		decodeRes, err := DecodeToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return ctx.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			})
		}

		_, err = repository.NewAuthRepository(config.NewRedisClient()).GetToken(context.Background(), decodeRes.AccessUuid)
		if err != nil {
			exception.PanicIfNeeded(errors.New(model.UNAUTHORIZATION))
		}

		//set global variable
		ctx.Set("currentId", decodeRes.UserId)
		ctx.Set("currentUsername", decodeRes.Username)
		ctx.Set("currentAccessUuid", decodeRes.AccessUuid)

		return next(ctx)
	}
}
