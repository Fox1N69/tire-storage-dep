package database

import (
	"api/api/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "user=postgres password=8008 dbname=makar-deplom port=5432 sslmode=disable"
var DB *gorm.DB

func Connection() {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Databases connect...")
	}
	DB.AutoMigrate(&models.User{})
}
