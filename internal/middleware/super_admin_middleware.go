package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/utils/response"
	jwttoken "github.com/phn00dev/go-URL-Shortener/pkg/jwtToken"
)

func SuperAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "error", "Authorization header required")
			c.Abort()
			return
		}

		// Bearer token'ı al
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Error(c, http.StatusUnauthorized, "error", "Bearer token required")
			c.Abort()
			return
		}

		// Token'ı doğrula
		claims, err := jwttoken.ValidateAdminToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "error", "Invalid token")
			c.Abort()
			return
		}

		// Admin rolü kontrol et
		if claims.AdminRole != "super_admin" {
			response.Error(c, http.StatusUnauthorized, "error", "Access denied")
			c.Abort()
			return
		}

		// Kullanıcı bilgilerini context'e ekle
		c.Set("id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("admin_role", claims.AdminRole)
		c.Set("email", claims.Email)

		// Super admin rolü varsa işlemi devam ettir
		c.Next()
	}
}
