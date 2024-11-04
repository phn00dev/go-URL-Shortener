package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/service"
)

type adminHandlerImp struct {
	adminService service.AdminService
	validator    *validator.Validate
}

func NewAdminHandler(service service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: service,
		validator:    validator.New(),
	}
}

func (adminHandler adminHandlerImp) GetOneById(c *gin.Context) {
	adminIdStr := c.Param("adminId")
	adminId, _ := strconv.Atoi(adminIdStr)
	admin, err := adminHandler.adminService.FindOneById(adminId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (adminHandler adminHandlerImp) GetAll(c *gin.Context) {
	admins, err := adminHandler.adminService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch admins"})
		return
	}
	c.JSON(http.StatusOK, admins)
}

func (adminHandler adminHandlerImp) Create(c *gin.Context) {
	var createRequest dto.CreateAdminRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate createRequest
	if err := adminHandler.validator.Struct(createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := adminHandler.adminService.Create(createRequest)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Admin created successfully"})
}

func (adminHandler adminHandlerImp) Update(c *gin.Context) {
	adminIdStr := c.Param("adminId")
	adminId, _ := strconv.Atoi(adminIdStr)
	var updateRequest dto.UpdateAdminRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate updateRequest
	if err := adminHandler.validator.Struct(updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := adminHandler.adminService.Update(adminId, updateRequest)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin updated successfully"})
}
func (adminHandler adminHandlerImp) Delete(c *gin.Context) {
	adminIdStr := c.Param("adminId")
	adminId, _ := strconv.Atoi(adminIdStr)
	err := adminHandler.adminService.Delete(adminId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
