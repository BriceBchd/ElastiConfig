package main

import (
	"api/controllers"
	"api/middlewares"
	"api/models"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	// Connect to MongoDB
	models.ConnectDB()

	router := gin.Default()
	router.Use(apmgin.Middleware(router))

	api := router.Group("/api")
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	admin := router.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddleware())
	admin.GET("/user", controllers.CurrentUser)

	router.Run(":8080")
}
