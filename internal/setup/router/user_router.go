package router

import (
	"github.com/gin-gonic/gin"

	userConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/user/constructor"
)

func UserRoutes(r *gin.Engine) {
	userApiRoute := r.Group("/user/v1/api")
	{
		authRoute := userApiRoute.Group("/auth")
		{
			authRoute.POST("/login", userConstructor.UserHandler.LoginUser)
		}
	}
}
