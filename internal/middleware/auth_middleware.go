package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	jwttoken "github.com/phn00dev/go-URL-Shortener/pkg/jwtToken"

)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}
		claims, err := jwttoken.ValidateAdminToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Next()
	}
}
