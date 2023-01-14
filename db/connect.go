package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	// database main config
	Dialector postgres.Config

	// database second or additional configuration
	Options gorm.Config

	// DB
	Sql *gorm.DB

	// error
	Error error
}

type Pool struct {
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	SetMaxIdleConns uint8

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	SetMaxOpenConns uint8

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	SetConnMaxLifetime int64
}

func NewDB() (*gorm.DB, error) {
	dsn := "host=localhost user=aliffatulmf password=aliffatulmf dbname=stki sslmode=disable TimeZone=Asia/Jakarta"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
