package service

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: repo,
	}
}

func (adminService adminServiceImp) FindOneById(adminId int) (*model.Admin, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) FindAll() ([]model.Admin, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) Create(createRequst dto.CreateAdminRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) Update(adminId int, updateRequest dto.UpdateAminRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) Delete(adminId int) error {
	panic("admin service imp")
}
