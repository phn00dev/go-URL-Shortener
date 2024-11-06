package service

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type UserService interface {
	FindAll() ([]model.User, error)
	FindOne(userId int) (*model.User, error)
	Delete(userId int) error
	// user
	RegisterUser(registerRequest dto.RegisterUserRequest) error
	LoginUser(loginRequest dto.UserLoginRequest) (*dto.UserLoginResponse, error)
	GetUserById(userId int) (*model.User, error)
	Update(userId int, updateRequest dto.UpdateUserRequest) error
	UpdateUserPassword(userId int, updatePasswordRequest dto.UpdateUserPassword) error
}
