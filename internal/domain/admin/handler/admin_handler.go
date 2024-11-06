package handler

import "github.com/gin-gonic/gin"

type AdminHandler interface {
	GetOneById(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	// admin login
	LoginAdmin(c *gin.Context)
	// update admin data for login admin
	UpdataAdminData(c *gin.Context)
	UpdateAdminPassword(c *gin.Context)
}
