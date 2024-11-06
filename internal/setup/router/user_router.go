package router

import (
	"github.com/gin-gonic/gin"

	urlConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/url/constructor"
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
			// url routes
			userUrlRoute := userRoute.Group("/urls")
			// userin doreden ahli url-leri get edip almaly
			userUrlRoute.GET("/", urlConstructor.UrlHandler.GetAllUserUrls)
			userUrlRoute.GET("/:urlId", urlConstructor.UrlHandler.GetOneUserUrl)
			userUrlRoute.POST("/create", urlConstructor.UrlHandler.Create)
			userUrlRoute.DELETE("/:urlId", urlConstructor.UrlHandler.Delete)
		}

	}
	r.GET("/:shortUrl", urlConstructor.UrlHandler.RedirectToOriginalUrl)
}
