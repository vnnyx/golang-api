package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/vnnyx/golang-api/config"
	"github.com/vnnyx/golang-api/exception"
	"github.com/vnnyx/golang-api/model"
	"time"
)

func CreateToken(request model.JwtPayload, configuration *config.Config) *model.TokenDetails {
	accessExpired := configuration.JWTMinute

	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessExpired)).Unix()
	td.AccessUuid = uuid.New().String()

	keyAccess, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(configuration.JWTSecretKey))
	exception.PanicIfNeeded(err)

	now := time.Now().UTC()

	atClaims := jwt.MapClaims{}
	atClaims["id"] = request.UserId
	atClaims["username"] = request.Username
	atClaims["email"] = request.Email
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["exp"] = td.AtExpires
	atClaims["iat"] = now.Unix()
	atClaims["iss"] = "simple-api"
	atClaims["aud"] = "simple-api"

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at.Header["simple-api"] = "jwt"
	td.AccessToken, err = at.SignedString(keyAccess)

	if err != nil {
		exception.PanicIfNeeded(errors.New(model.AUTHENTICATION_FAILURE_ERR_TYPE))
	}

	return td
}
