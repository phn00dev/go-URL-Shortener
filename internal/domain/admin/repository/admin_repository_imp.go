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
	var admin model.Admin
	err := adminRepo.db.First(&admin, adminId).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (adminRepo adminRepositoryImp) GetAll() ([]model.Admin, error) {
	var admins []model.Admin
	if err := adminRepo.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (adminRepo adminRepositoryImp) Create(admin model.Admin) error {
	return adminRepo.db.Create(&admin).Error
}

func (adminRepo adminRepositoryImp) Update(adminId int, admin model.Admin) error {
	return adminRepo.db.Where("id=?", adminId).Updates(&admin).Error
}

func (adminRepo adminRepositoryImp) Delete(adminId int) error {
	return adminRepo.db.Delete(&model.Admin{}, adminId).Error
}

func (adminRepo adminRepositoryImp) GetAdminByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := adminRepo.db.Where("email=?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (adminRepo adminRepositoryImp) GetAdminByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := adminRepo.db.Where("username=?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (adminRepo adminRepositoryImp) UpdateAdminPassword(adminId int, password string) error {
	if err := adminRepo.db.Model(&model.Admin{}).Where("id = ?", adminId).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
