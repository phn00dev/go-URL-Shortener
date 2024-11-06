package repository

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type UserRepository interface {
	GetById(userId int) (*model.User, error)
	GetAll() ([]dto.AllUserResponse, error)
	Delete(userId int) error

	// user
	Create(user model.User) error
	Update(userId int, user model.User) error
	UpdateUserPassword(userId int, password string) error
	GetByEmail(email string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	FindByUsernameOrEmail(username, email string) (*model.User, error)
	FindByUsernameOrEmailById(userId int, username, email string) (*model.User, error)
}
