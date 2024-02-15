package database

import (
	"api/api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
}
