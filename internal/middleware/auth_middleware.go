package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/utils/response"
	jwttoken "github.com/phn00dev/go-URL-Shortener/pkg/jwtToken"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "error", "Authorization header required")
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Error(c, http.StatusUnauthorized, "error", "Bearer token required")
			c.Abort()
			return
		}
		claims, err := jwttoken.ValidateToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "error", "Invalid token")
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Next()
	}
}
