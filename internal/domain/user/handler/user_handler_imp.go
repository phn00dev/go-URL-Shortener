package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/service"
	"github.com/phn00dev/go-URL-Shortener/internal/utils/response"
	"github.com/phn00dev/go-URL-Shortener/internal/utils/validate"

)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}
}

func (userHandler userHandlerImp) GetById(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	user, err := userHandler.userService.FindOne(userId)
	if err != nil {
		response.Error(c, http.StatusNotFound, "user Not Found", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "User retrieved successfully", user)
}

func (userHandler userHandlerImp) GetAll(c *gin.Context) {
	users, err := userHandler.userService.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "users not found", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "users retrieved successfully", users)
}

func (userHandler userHandlerImp) Create(c *gin.Context) {
	var createRequest dto.CreateUserRequest
	if err := c.ShouldBindBodyWithJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}
	// validate data
	if err := validate.ValidateStruct(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}

	// create user

	if err := userHandler.userService.Create(createRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "create user error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user created successfully", nil)
}

func (userHandler userHandlerImp) Update(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	var updateRequest dto.UpdateUserRequest
	if err := c.ShouldBindBodyWithJSON(&updateRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}

	// validate error
	if err := validate.ValidateStruct(updateRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}
	// update user data
	if err := userHandler.userService.Update(userId, updateRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "update user error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user updated successfully", nil)
}

func (userHandler userHandlerImp) Delete(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	if err := userHandler.userService.Delete(userId); err != nil {
		response.Error(c, http.StatusInternalServerError, "user deleted error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user deleted successfully", nil)
}

func (userHandler userHandlerImp) UpdateUserPassword(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	var updatePasswordRequest dto.UpdateUserPassword
	if err := c.ShouldBindBodyWithJSON(&updatePasswordRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}

	// validate error
	if err := validate.ValidateStruct(updatePasswordRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}

	// update user password
	if err := userHandler.userService.UpdateUserPassword(userId, updatePasswordRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "update password error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user password updated successfully", nil)
}

func (userHandler userHandlerImp) LoginUser(c *gin.Context) {

	var userLoginRequest dto.UserLoginRequest
	if err := c.ShouldBindBodyWithJSON(&userLoginRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}

	// validate error
	if err := validate.ValidateStruct(userLoginRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}

	// login service

	loginResponse, err := userHandler.userService.LoginUser(userLoginRequest)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user login successfully", loginResponse)
}
