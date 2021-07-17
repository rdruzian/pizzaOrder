package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pizzaOrder/database/migrations"
	"time"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=pizzaOrder sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
