package router

import "github.com/gin-gonic/gin"

func AdminRoutes(r *gin.Engine) {
	r.GET("/admin", func(c *gin.Context) {
		// Burada koduňyzy ýazyň
		c.JSON(200, gin.H{
			"message": "Admin Panel",
		})
	})
}
