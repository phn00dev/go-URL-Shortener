package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/service"
	"github.com/phn00dev/go-URL-Shortener/internal/utils/response"
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
	adminId, err := strconv.Atoi(adminIdStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid admin ID", err.Error())
		return
	}

	admin, err := adminHandler.adminService.FindOneById(adminId)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Admin not found", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Admin retrieved successfully", admin)
}

func (adminHandler adminHandlerImp) GetAll(c *gin.Context) {
	admins, err := adminHandler.adminService.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Could not fetch admins", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Admins retrieved successfully", admins)
}

func (adminHandler adminHandlerImp) Create(c *gin.Context) {
	var createRequest dto.CreateAdminRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "data parsing error", err.Error())
		return
	}

	// Validate createRequest
	if err := adminHandler.validator.Struct(createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	err := adminHandler.adminService.Create(createRequest)
	if err != nil {
		response.Error(c, http.StatusConflict, "admin creation error", err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Admin created successfully", nil)
}

func (adminHandler adminHandlerImp) Update(c *gin.Context) {
	adminIdStr := c.Param("adminId")
	adminId, err := strconv.Atoi(adminIdStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid admin ID", err.Error())
		return
	}

	var updateRequest dto.UpdateAdminRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "data parsing error", err.Error())
		return
	}

	// Validate updateRequest
	if err := adminHandler.validator.Struct(updateRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	err = adminHandler.adminService.Update(adminId, updateRequest)
	if err != nil {
		response.Error(c, http.StatusConflict, "admin update error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Admin updated successfully", nil)
}

func (adminHandler adminHandlerImp) Delete(c *gin.Context) {
	adminIdStr := c.Param("adminId")
	adminId, err := strconv.Atoi(adminIdStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid admin ID", err.Error())
		return
	}

	err = adminHandler.adminService.Delete(adminId)
	if err != nil {
		response.Error(c, http.StatusNotFound, "admin not found", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Admin deleted successfully", nil)
}
