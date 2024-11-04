package service

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type AdminService interface {
	FindOneById(adminId int) (*model.Admin, error)
	FindAll() ([]model.Admin, error)
	Create(createRequst dto.CreateAdminRequest) error
	Update(adminId int, updateRequest dto.UpdateAminRequest) error
	Delete(adminId int) error
}
