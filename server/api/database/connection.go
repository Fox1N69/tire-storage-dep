package database

import (
	"api/api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Migrator().CreateTable(&models.User{})
	db.Debug().AutoMigrate(&models.User{})
}
