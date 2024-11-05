package repository

import "github.com/phn00dev/go-URL-Shortener/internal/model"

type AdminRepository interface {
	GetOneById(adminId int) (*model.Admin, error)
	GetAll() ([]model.Admin, error)
	Create(admin model.Admin) error
	Update(adminId int, admin model.Admin) error
	Delete(adminId int) error
	UpdateAdminPassword(adminId int, password string) error
	// find
	FindByUsernameOrEmail(username, email string) (*model.Admin, error)
	FindByUsernameOrEmailById(adminId int, username, email string) (*model.Admin, error)
	GetAdminByEmail(email string) (*model.Admin, error)
	GetAdminByUsername(username string) (*model.Admin, error)
}
