package main

import (
	"fmt"
	"go-bookstore-api/config"
)

func main() {
	config.ConnectDB()

	// Check if the database connection is established
	if config.DB != nil {
		fmt.Println("Database connection established successfully.")
	} else {
		fmt.Println("Failed to establish database connection.")
	}
}
