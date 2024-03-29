package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vanbien2402/first-web-demo/internal/pkg/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		if err := jwt.ValidateToken(tokenString); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
