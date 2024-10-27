package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB инициализирует подключение к базе данных.
func InitDB() {
	var (
		err error
	)
	// Обновленный DSN
	dsn := "host=127.0.0.1 user=myuser password=mypassword dbname=mydatabase port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}
