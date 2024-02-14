package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() {
	if db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}
	
}
