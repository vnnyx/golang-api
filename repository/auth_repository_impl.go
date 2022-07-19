package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/vnnyx/golang-api/model"
	"time"
)

type AuthRepositoryImpl struct {
	Redis *redis.Client
}

func NewAuthRepository(redis *redis.Client) AuthRepository {
	return &AuthRepositoryImpl{Redis: redis}
}

func (repo *AuthRepositoryImpl) StoreToken(ctx context.Context, details model.TokenDetails) error {
	err := repo.Redis.Set(ctx, details.AccessUuid, details.AccessToken, time.Unix(details.AtExpires, 0).Sub(time.Now())).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthRepositoryImpl) DeleteToken(ctx context.Context, accessUuid string) error {
	err := repo.Redis.Del(ctx, accessUuid).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthRepositoryImpl) GetToken(ctx context.Context, accessUuid string) (res string, err error) {
	res, err = repo.Redis.Get(ctx, accessUuid).Result()
	if err != nil {
		return res, err
	}
	return res, nil
}
