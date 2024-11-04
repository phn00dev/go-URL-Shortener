package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type adminRepositoryImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepositoryImp{
		db: db,
	}
}

func (adminRepo adminRepositoryImp) GetOneById(adminId int) (*model.Admin, error) {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) GetAll() ([]model.Admin, error) {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Create(admin model.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Update(adminId int, admin model.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Delete(adminId int) error {
	panic("admin repo imp")
}
