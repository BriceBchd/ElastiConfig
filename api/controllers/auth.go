package controllers

import (
	"api/models"
	"api/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Login(c *gin.Context) {
	var input LoginInput

	// Verify input email and password are provided
	if err := c.ShouldBindJSON(&input); err != nil {
		// If no JSON body is provided, return an error
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no input provided"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user
	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password

	// Verify the user's credentials against a database
	token, err := models.LoginCheck(u.Email, u.Password)
	print(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	// Verify input email and password are provided
	if err := c.ShouldBindJSON(&input); err != nil {
		// If no JSON body is provided, return an error
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no input provided"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user
	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password

	// Save the user to the database
	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
