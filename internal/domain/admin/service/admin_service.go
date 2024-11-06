package service

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type AdminService interface {
	FindOneById(adminId int) (*model.Admin, error)
	FindAll() ([]model.Admin, error)
	Create(createRequst dto.CreateAdminRequest) error
	Update(adminId int, updateRequest dto.UpdateAdminRequest) error
	Delete(adminId int) error
	// login bolan adminin maglumatlaryny uytgetmek

	AdminLogin(loginRequest dto.AdminLoginRequest) (*dto.AdminLoginResponse, error)
	UpdateAdminPassword(adminId int, changePasswordRequest dto.ChangeAdminPassword) error
	UpdateAdminData(adminId int, updateAdminData dto.UpdateLoginAdminRequest) error
}
