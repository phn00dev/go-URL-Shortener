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

func (userHandler userHandlerImp) GetAll(c *gin.Context) {
	users, err := userHandler.userService.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "users not found", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "users retrieved successfully", users)
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

func (userHandler userHandlerImp) Delete(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	if err := userHandler.userService.Delete(userId); err != nil {
		response.Error(c, http.StatusInternalServerError, "user deleted error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user deleted successfully", nil)
}

// user routes
func (userHandler userHandlerImp) RegisterUser(c *gin.Context) {
	var registerRequest dto.RegisterUserRequest
	if err := c.ShouldBindBodyWithJSON(&registerRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "body parser error", err.Error())
		return
	}
	// validate data
	if err := validate.ValidateStruct(&registerRequest); err != nil {
		response.Error(c, http.StatusBadRequest, "validate error", err.Error())
		return
	}
	// create user
	if err := userHandler.userService.RegisterUser(registerRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "create user error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user  registered successfully", nil)
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

func (userHandler userHandlerImp) GetUser(c *gin.Context) {
	authUserId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authUserIdInt, ok := authUserId.(int)
	if !ok || authUserIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	// call user GetUserById method
	user, err := userHandler.userService.GetUserById(authUserIdInt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user account information", user)
}

func (userHandler userHandlerImp) UpdateUser(c *gin.Context) {
	authUserId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authUserIdInt, ok := authUserId.(int)
	if !ok || authUserIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

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
	if err := userHandler.userService.Update(authUserIdInt, updateRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "update user error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user updated successfully", nil)
}

func (userHandler userHandlerImp) UpdateUserPassword(c *gin.Context) {
	authUserId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authUserIdInt, ok := authUserId.(int)
	if !ok || authUserIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

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
	if err := userHandler.userService.UpdateUserPassword(authUserIdInt, updatePasswordRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "update password error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user password updated successfully", nil)
}

func (userHandler userHandlerImp) DeleteProfile(c *gin.Context) {
	authUserId, exists := c.Get("id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}

	authUserIdInt, ok := authUserId.(int)
	if !ok || authUserIdInt == 0 {
		response.Error(c, http.StatusUnauthorized, "error auth", "User not authorized")
		return
	}
	if err := userHandler.userService.Delete(authUserIdInt); err != nil {
		response.Error(c, http.StatusInternalServerError, "user deleted error", err.Error())
		return
	}
	response.Success(c, http.StatusOK, "user profile deleted successfully", nil)
}
