package handler

import "github.com/gin-gonic/gin"

type UserRepository interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UpdateUserPassword(c *gin.Context)
}
