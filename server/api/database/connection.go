package database

import (
	"api/api/models"
	"log"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var dsn = "user=postgres password=8008 dbname=deplom-makar sslmode=disable"
var DB *gorm.DB

func Connection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(&models.User{})
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = Connection()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			log.Print("database is unavaibel. Wait for %d sec.\n")
			time.Sleep(sleep * time.Second)
			DB = Connection()
		}
	}
	return DB
}
