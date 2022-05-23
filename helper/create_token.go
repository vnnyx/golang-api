package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang-simple-api/exception"
	"golang-simple-api/model"
	"os"
	"strconv"
	"time"
)

func CreateToken(request model.JwtPayload) *model.TokenDetails {
	accessExpired, _ := strconv.Atoi(os.Getenv("JWT_MINUTE"))

	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessExpired)).Unix()
	td.AccessUuid = uuid.New().String()

	var err error

	keyAccess, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_SECRET_KEY")))
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
