package router

import (
	"github.com/gin-gonic/gin"

	adminConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/admin/constructor"
)

func AdminRoutes(route *gin.Engine) {
	adminApiRoute := route.Group("/api/v1/admin")
	adminRoute := adminApiRoute.Group("/admins")
	{
		adminRoute.GET("/", adminConstructor.AdminHandler.GetAll)
		adminRoute.GET("/:adminId", adminConstructor.AdminHandler.GetOneById)
		adminRoute.POST("/create", adminConstructor.AdminHandler.Create)
		adminRoute.PUT("/:adminId", adminConstructor.AdminHandler.Update)
		adminRoute.DELETE("/:adminId", adminConstructor.AdminHandler.Delete)
	}
}
