package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type userRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		db: db,
	}
}

func (userRepo userRepositoryImp) GetById(userId int) (*model.User, error) {
	var user model.User
	if err := userRepo.db.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := userRepo.db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := userRepo.db.Where("username=?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) GetAll() ([]model.User, error) {
	var users []model.User
	if err := userRepo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepo userRepositoryImp) Create(user model.User) error {
	return userRepo.db.Create(&user).Error
}

func (userRepo userRepositoryImp) Update(userId int, user model.User) error {
	return userRepo.db.Where("id=?", userId).Updates(&user).Error
}

func (userRepo userRepositoryImp) Delete(userId int) error {
	return userRepo.db.Delete(&model.User{}, userId).Error
}

func (userRepo userRepositoryImp) UpdateUserPassword(userId int, password string) error {
	if err := userRepo.db.Model(&model.User{}).Where("id = ?", userId).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}