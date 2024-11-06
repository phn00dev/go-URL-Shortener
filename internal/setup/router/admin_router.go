package router

import (
	"github.com/gin-gonic/gin"

	adminConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/admin/constructor"
	urlConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/url/constructor"
	userConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/user/constructor"
	"github.com/phn00dev/go-URL-Shortener/internal/middleware"
)

func AdminRoutes(route *gin.Engine) {
	adminApiRoute := route.Group("/v1a/api/admin")
	{
		// admin auth routes
		authAdminRoute := adminApiRoute.Group("/auth")
		{
			authAdminRoute.POST("/login", adminConstructor.AdminHandler.LoginAdmin)
		}
		// admin routes
		adminApiRoute.Use(middleware.AuthMiddleware())
		adminRoute := adminApiRoute.Group("/admins")
		{
			adminRoute.GET("/", adminConstructor.AdminHandler.GetAll)
			adminRoute.GET("/:adminId", adminConstructor.AdminHandler.GetOneById)
			adminRoute.POST("/create", adminConstructor.AdminHandler.Create)
			adminRoute.PUT("/:adminId", adminConstructor.AdminHandler.Update)
			adminRoute.DELETE("/:adminId", adminConstructor.AdminHandler.Delete)
		}
		// user routes
		userRoute := adminApiRoute.Group("/users")
		{
			userRoute.GET("/", userConstructor.UserHandler.GetAll)
			userRoute.GET("/:userId", userConstructor.UserHandler.GetById)
			// userRoute.PUT("/:userId", userConstructor.UserHandler.Update)
			userRoute.DELETE("/:userId", userConstructor.UserHandler.Delete)
		}
		// url routes
		urlRoute := adminApiRoute.Group("/urls")
		{
			urlRoute.GET("/", urlConstructor.UrlHandler.GetAll)
			urlRoute.GET("/:urlId", urlConstructor.UrlHandler.GetOne)
		}
	}

}
