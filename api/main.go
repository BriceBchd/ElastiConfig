package main

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)

	router.Run(":8080")
}
