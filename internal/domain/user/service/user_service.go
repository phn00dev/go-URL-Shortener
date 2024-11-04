package service

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type UserService interface {
	FindAll() ([]model.User, error)
	FindOne(userId int) (*model.User, error)
	Create(createRequest dto.CreateUserRequest) error
	Update(userId int, updateRequest dto.UpdateUserRequest) error
	Delete(userId int) error
}
