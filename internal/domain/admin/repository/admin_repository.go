package repository

import "github.com/phn00dev/go-URL-Shortener/internal/model"

type AdminRepository interface {
	GetOneById(adminId int) (*model.Admin, error)
	GetAll() ([]model.Admin, error)
	Create(admin model.Admin) error
	Update(adminId int, admin model.Admin) error
	Delete(adminId int) error
}
