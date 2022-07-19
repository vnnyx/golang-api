package repository

import (
	"context"
	"github.com/vnnyx/golang-api/model"
)

type AuthRepository interface {
	StoreToken(ctx context.Context, details model.TokenDetails) error
	DeleteToken(ctx context.Context, accessUuid string) error
	GetToken(ctx context.Context, accessUuid string) (res string, err error)
}
