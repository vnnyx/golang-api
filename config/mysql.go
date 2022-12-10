package config

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vnnyx/golang-api/exception"
	"time"
)

func NewMySQLDatabase(configuration *Config) *sqlx.DB {
	ctx, cancel := NewMySQLContext()
	defer cancel()

	sqlDB, err := sqlx.Open("mysql", configuration.MysqlHostSlave)
	exception.PanicIfNeeded(err)

	err = sqlDB.PingContext(ctx)
	exception.PanicIfNeeded(err)

	mysqlPoolMax := configuration.MysqlPoolMax

	mysqlIdleMax := configuration.MysqlIdleMax

	mysqlMaxLifeTime := configuration.MysqlMaxLifeTimeMinute

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(mysqlIdleMax)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(mysqlPoolMax)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlMaxLifeTime) * time.Minute)

	//sqlDB.SetConnMaxIdleTime(time.Duration(mysqlMaxIdleTime) * time.Minute)

	return sqlDB
}

func NewMySQLContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
