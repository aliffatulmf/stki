package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	dsn := "host=172.24.10.172 user=aliffatulmf password=aliffatulmf dbname=stki sslmode=disable TimeZone=Asia/Jakarta"
	dialector := postgres.New(postgres.Config{
		DSN: dsn,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
