package handlers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandler is a handler for the /register route
func RegisterHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Store the user in a database

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// LoginHandler is a handler for the /login route
func LoginHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Verify the user's credentials against a database

	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}
