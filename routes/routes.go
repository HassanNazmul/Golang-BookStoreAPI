package routes

import (
	"go-bookstore-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Book Routes
	r.GET("/books", controllers.GetBooks) // Fetch all books
	//r.GET("/books/:id", controllers.GetBook)       // Fetch a book by ID
	//r.POST("/books", controllers.CreateBook)       // Create a new book
	//r.PUT("/books/:id", controllers.UpdateBook)    // Update a book
	//r.DELETE("/books/:id", controllers.DeleteBook) // Delete a book

	return r
}
