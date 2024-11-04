package app

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/setup/router"
	"github.com/phn00dev/go-URL-Shortener/pkg/config"
)

func NewApp(config *config.Config) (httpServer *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	// Default gin.Engine döretmek
	httpServer = gin.New()

	// Middleware goşmak
	httpServer.Use(gin.Logger())
	httpServer.Use(gin.Recovery())

	// Konfigurasiýalary geçirmegiň beýany
	httpServer.Use(func(c *gin.Context) {
		// Eger hata bolsa, JSON döndür
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{
					"status":  500,
					"message": "Näsazlyk ýüze çykdy, Sonrak synanysyn!!!",
				})
			}
		}()
		c.Next()
	})

	// HTTP server konfigurasiýalary
	httpServer.Use(func(c *gin.Context) {
		c.Set("AppName", config.HttpConfig.AppName)        // App adyny geçir
		c.Set("ServerHeader", config.HttpConfig.AppHeader) // Server header geçir
		// Başga konfigurasiýalary geçiriň
	})
	httpServer.SetTrustedProxies([]string{"*"})
	// Routerlary goşmak
	router.AdminRoutes(httpServer)
	router.UserRoutes(httpServer)

	return httpServer
}
