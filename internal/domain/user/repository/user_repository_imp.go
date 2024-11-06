package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/dto"
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
	if err := userRepo.db.
		Table("users").
		Select("users.id, users.username, users.email, TO_CHAR(users.created_at, 'YYYY-MM-DD HH24:MI:SS') as created_at, COALESCE(COUNT(urls.id), 0) AS url_count").
		Joins("LEFT JOIN urls ON urls.user_id = users.id").
		Group("users.id").
		Preload("UserUrls").
		First(&user, userId).Error; err != nil {
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

func (userRepo userRepositoryImp) GetAll() ([]dto.AllUserResponse, error) {
	var users []dto.AllUserResponse

	if err := userRepo.db.
		Table("users").
		Select("users.id, users.username, users.email, TO_CHAR(users.created_at, 'YYYY-MM-DD HH24:MI:SS') as created_at, COALESCE(COUNT(urls.id), 0) AS url_count").
		Joins("LEFT JOIN urls ON urls.user_id = users.id").
		Group("users.id").
		Order("users.id desc").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// user

func (userRepo userRepositoryImp) Delete(userId int) error {
	return userRepo.db.Delete(&model.User{}, userId).Error
}

func (userRepo userRepositoryImp) Create(user model.User) error {
	return userRepo.db.Create(&user).Error
}

func (userRepo userRepositoryImp) Update(userId int, user model.User) error {
	return userRepo.db.Where("id=?", userId).Updates(&user).Error
}

func (userRepo userRepositoryImp) UpdateUserPassword(userId int, password string) error {
	if err := userRepo.db.Model(&model.User{}).Where("id = ?", userId).Update("password_hash", password).Error; err != nil {
		return err
	}
	return nil
}

func (userRepo userRepositoryImp) FindByUsernameOrEmail(username, email string) (*model.User, error) {
	var user model.User
	if err := userRepo.db.Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Gaýtalanmaýan ýagdaýda nil gaýtarýar
		}
		return nil, err // Başga ýalňyşlyk bar bolsa, ýalňyşlygy gaýtarýar
	}
	return &user, nil
}

func (userRepo userRepositoryImp) FindByUsernameOrEmailById(userId int, username, email string) (*model.User, error) {
	var user model.User
	if err := userRepo.db.Where("id !=?", userId).Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
