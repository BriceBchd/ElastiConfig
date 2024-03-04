package main

import (
	"api/controllers"
	"api/middlewares"
	"api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	models.ConnectDB()

	router := gin.Default()

	api := router.Group("/api")
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	admin := router.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddleware())
	admin.GET("/user", controllers.CurrentUser)

	router.Run(":8080")
}
