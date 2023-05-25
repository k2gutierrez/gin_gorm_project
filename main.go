package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k2gutierrez/gin_gorm_project/controllers"
	"github.com/k2gutierrez/gin_gorm_project/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()

}
