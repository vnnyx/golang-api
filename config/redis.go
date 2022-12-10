package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/vnnyx/golang-api/exception"
)

func NewRedisClient() *redis.Client {
	config, err := NewConfig(".", ".env")
	exception.PanicIfNeeded(err)
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword,
	})
	return client
}
