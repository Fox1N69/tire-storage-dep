package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() {
	if _, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}); err != nil {
		log.Fatal(err, "Error connection to database")
	}
}
