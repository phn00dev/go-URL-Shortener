package router

import (
	"github.com/gin-gonic/gin"

	adminConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/admin/constructor"
	userConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/user/constructor"
)

func AdminRoutes(route *gin.Engine) {
	adminApiRoute := route.Group("/v1a/api/admin")
	adminRoute := adminApiRoute.Group("/admins")
	{
		adminRoute.GET("/", adminConstructor.AdminHandler.GetAll)
		adminRoute.GET("/:adminId", adminConstructor.AdminHandler.GetOneById)
		adminRoute.POST("/create", adminConstructor.AdminHandler.Create)
		adminRoute.PUT("/:adminId", adminConstructor.AdminHandler.Update)
		adminRoute.DELETE("/:adminId", adminConstructor.AdminHandler.Delete)
	}

	userRoute := adminApiRoute.Group("/users")
	{
		userRoute.GET("/", userConstructor.UserHandler.GetAll)
		userRoute.GET("/:userId", userConstructor.UserHandler.GetById)
		userRoute.POST("/create", userConstructor.UserHandler.Create)
		userRoute.PUT("/:userId", userConstructor.UserHandler.Update)
		userRoute.DELETE("/:userId", userConstructor.UserHandler.Delete)
	}

}
