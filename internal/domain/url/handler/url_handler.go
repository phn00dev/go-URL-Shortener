package handler

import "github.com/gin-gonic/gin"

type UrlHandler interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	// user urls
	GetAllUserUrls(c *gin.Context)
	GetOneUserUrl(c *gin.Context)
}
