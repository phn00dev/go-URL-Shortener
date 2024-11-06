package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Delete(c *gin.Context)

	// user profile
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateUserPassword(c *gin.Context)
	DeleteProfile(c *gin.Context)
}
