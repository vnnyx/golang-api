package config

import (
	"github.com/spf13/viper"
	"github.com/vnnyx/golang-api/exception"
)

type Config struct {
	MysqlHostSlave         string `mapstructure:"MYSQL_HOST_SLAVE"`
	MysqlPoolMin           int    `mapstructure:"MYSQL_POOL_MIN"`
	MysqlPoolMax           int    `mapstructure:"MYSQL_POOL_MAX"`
	MysqlIdleMax           int    `mapstructure:"MYSQL_IDLE_MAX"`
	MysqlMaxIdleTimeMinute int    `mapstructure:"MYSQL_MAX_IDLE_TIME_MINUTE"`
	MysqlMaxLifeTimeMinute int    `mapstructure:"MYSQL_MAX_LIFE_TIME_MINUTE"`
	JWTPublicKey           string `mapstructure:"JWT_PUBLIC_KEY"`
	JWTSecretKey           string `mapstructure:"JWT_SECRET_KEY"`
	JWTMinute              int    `mapstructure:"JWT_MINUTE"`
	RedisHost              string `mapstructure:"REDIS_HOST"`
	RedisPassword          string `mapstructure:"REDIS_PASSWORD"`
}

func NewConfig(path string, configName string) (*Config, error) {
	config := &Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	exception.PanicIfNeeded(err)
	err = viper.Unmarshal(&config)
	exception.PanicIfNeeded(err)
	return config, err
}
