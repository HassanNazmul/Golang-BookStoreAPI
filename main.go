package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bookstore-api/config"
	"go-bookstore-api/routes"
)

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.DebugMode)
	fmt.Println("Gin Mode:", gin.Mode())

	// Connect to the database
	config.ConnectDB()

	// Set up the routes
	r := routes.SetupRoutes()

	fmt.Println("Starting server on port 8080")
	if err := r.Run(":8080"); err != nil { // Start the server on port 8080
		fmt.Printf("Failed to start server: %v\n", err)
	} else {
		fmt.Println("Server started successfully")
	}
}
