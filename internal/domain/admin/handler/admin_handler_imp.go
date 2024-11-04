package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/service"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(service service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: service,
	}
}

func (adminHandler adminHandlerImp) GetOneById(c *gin.Context) {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) GetAll(c *gin.Context) {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Create(c *gin.Context) {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Update(c *gin.Context) {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Delete(c *gin.Context) {
	panic("admin handler imp")
}
