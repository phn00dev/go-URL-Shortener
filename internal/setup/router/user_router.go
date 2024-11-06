package router

import (
	"github.com/gin-gonic/gin"

	userConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/user/constructor"
	"github.com/phn00dev/go-URL-Shortener/internal/middleware"

)

func UserRoutes(r *gin.Engine) {
	userApiRoute := r.Group("/user/v1/api")
	{
		authRoute := userApiRoute.Group("/auth")
		{
			authRoute.POST("/register", userConstructor.UserHandler.RegisterUser)
			authRoute.POST("/login", userConstructor.UserHandler.LoginUser)
		}
		// user routes

		userRoute := userApiRoute.Group("/user")
		userRoute.Use(middleware.AuthMiddleware())
		{
			userRoute.GET("/", userConstructor.UserHandler.GetUser)
			userRoute.PUT("/update", userConstructor.UserHandler.UpdateUser)
			userRoute.PUT("/change-password", userConstructor.UserHandler.UpdateUserPassword)
			userRoute.DELETE("/delete-profile", userConstructor.UserHandler.DeleteProfile)
		}

	}
}
