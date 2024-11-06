package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/service"
	"github.com/phn00dev/go-URL-Shortener/internal/utils/response"
	"github.com/phn00dev/go-URL-Shortener/internal/utils/validate"
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
	if err := c.ShouldBindBodyWithJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
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
	if err := c.ShouldBindBodyWithJSON(&updateRequest); err != nil {
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

func (adminHandler adminHandlerImp) LoginAdmin(c *gin.Context) {
	var adminLoginRequest dto.AdminLoginRequest
	// Pointer görnüşinde geçiriň: &adminLoginRequest
	if err := c.ShouldBindJSON(&adminLoginRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}
	// validate data
	if err := validate.ValidateStruct(adminLoginRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}
	// call login service
	loginResponse, err := adminHandler.adminService.AdminLogin(adminLoginRequest)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "admin login successfully", loginResponse)
}

// update admin data for login admin

func (adminHandler adminHandlerImp) UpdataAdminData(c *gin.Context) {
	authAdminId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authAdminIdInt, ok := authAdminId.(int)
	if !ok || authAdminIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}
	var updateAdminData dto.UpdateLoginAdminRequest
	// body parser
	if err := c.ShouldBindBodyWithJSON(&updateAdminData); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}

	// validate data
	if err := validate.ValidateStruct(updateAdminData); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}

	// update admin data
	if err := adminHandler.adminService.UpdateAdminData(authAdminIdInt, updateAdminData); err != nil {
		response.Error(c, http.StatusInternalServerError, "update admin data error", err.Error())
		return
	}
	// success respponse
	response.Success(c, http.StatusOK, "admin data updated successfully", nil)
}

func (adminHandler adminHandlerImp) UpdateAdminPassword(c *gin.Context) {
	authAdminId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authAdminIdInt, ok := authAdminId.(int)
	if !ok || authAdminIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	var changeAdminPassword dto.ChangeAdminPassword
	// body parser
	if err := c.ShouldBindBodyWithJSON(&changeAdminPassword); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}

	// validate data
	if err := validate.ValidateStruct(changeAdminPassword); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}
	// change admin password service
	if err := adminHandler.adminService.UpdateAdminPassword(authAdminIdInt, changeAdminPassword); err != nil {
		response.Error(c, http.StatusInternalServerError, "update password error", err.Error())
		return
	}
	// success respponse
	response.Success(c, http.StatusOK, "admin password updated successfully", nil)
}
