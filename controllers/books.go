package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k2gutierrez/gin_gorm_project/models"
)

// GET /books
// Get all the books

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
