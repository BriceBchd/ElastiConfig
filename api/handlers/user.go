package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Store the user in a database

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func LoginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Verify the user's credentials against a database

	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}
