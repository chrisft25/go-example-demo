package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection(dsn string){
	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal("Failed to connect to database!")
	} else {
		log.Println("Connected to database!")
	}
}