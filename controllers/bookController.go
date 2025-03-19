package controllers

import (
	"go-bookstore-api/config"
	"go-bookstore-api/models"
	"net/http" // Import the http package

	"github.com/gin-gonic/gin"
)

// GetBooks - Fetch all books
func GetBooks(c *gin.Context) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
