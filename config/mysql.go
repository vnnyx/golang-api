package config

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang-simple-api/exception"
	"strconv"
	"time"
)

func NewMySQLDatabase(configuration Config) *sql.DB {
	ctx, cancel := NewMySQLContext()
	defer cancel()

	sqlDB, err := sql.Open("mysql", configuration.Get("MYSQL_HOST_SLAVE"))
	exception.PanicIfNeeded(err)

	err = sqlDB.PingContext(ctx)
	exception.PanicIfNeeded(err)

	mysqlPoolMax, err := strconv.Atoi(configuration.Get("MYSQL_POOL_MAX"))
	exception.PanicIfNeeded(err)

	mysqlIdleMax, err := strconv.Atoi(configuration.Get("MYSQL_IDLE_MAX"))
	exception.PanicIfNeeded(err)

	mysqlMaxLifeTime, err := strconv.Atoi(configuration.Get("MYSQL_MAX_LIFE_TIME_MINUTE"))
	exception.PanicIfNeeded(err)

	// mysqlMaxIdleTime, err := strconv.Atoi(configuration.Get("MYSQL_MAX_IDLE_TIME_MINUTE"))
	exception.PanicIfNeeded(err)

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
