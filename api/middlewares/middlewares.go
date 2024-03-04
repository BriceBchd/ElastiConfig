package middlewares

import (
	"net/http"

	"api/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.String(http.StatusUnauthorized, "Unauthorized, no token provided")
			c.Abort()
			return
		}
		_, err := token.TokenValid(tokenString)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
