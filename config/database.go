package config

import (
	"log"

	// Import the Book model
	"go-bookstore-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the PostgreSQL database using GORM.

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres dbname=bookstore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Explicitly migrate the schema
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Error during schema migration: %v", err)
	}

	DB = db
	log.Println("Database Connected successfully")
}
