package service

import (
	"errors"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
	"github.com/phn00dev/go-URL-Shortener/internal/utils"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (s userServiceImp) FindAll() ([]model.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s userServiceImp) FindOne(userId int) (*model.User, error) {
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s userServiceImp) Create(createRequest dto.CreateUserRequest) error {
	email := createRequest.Email
	username := createRequest.Username
	if err := validateUniqueEmailAndUsername(email, username, s.userRepo); err != nil {
		return err
	}
	newUser := model.User{
		Email:        email,
		Username:     username,
		PasswordHash: utils.HashPassword(createRequest.Password),
	}
	return s.userRepo.Create(newUser)
}

func (s userServiceImp) Update(userId int, updateRequest dto.UpdateUserRequest) error {
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if updateRequest.Username != "" {
		if err := validateUniqueUsername(updateRequest.Username, userId, s.userRepo); err != nil {
			return err
		}
		user.Username = updateRequest.Username
	}
	if updateRequest.Email != "" {
		if err := validateUniqueEmail(updateRequest.Email, userId, s.userRepo); err != nil {
			return err
		}
		user.Email = updateRequest.Email
	}
	return s.userRepo.Update(user.ID, *user)
}

func (s userServiceImp) Delete(userId int) error {
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		return err
	}
	if user.ID != 0 {
		return errors.New("something error")
	}
	return s.userRepo.Delete(user.ID)
}

func validateUniqueEmail(email string, adminId int, repo repository.UserRepository) error {
	existingAdmin, err := repo.GetByEmail(email)
	if err != nil {
		return err
	}
	if existingAdmin != nil && existingAdmin.ID != adminId {
		return errors.New("this email is already used by another admin")
	}
	return nil
}

func validateUniqueUsername(username string, adminId int, repo repository.UserRepository) error {
	existingAdmin, err := repo.GetByUsername(username)
	if err != nil {
		return err
	}
	if existingAdmin != nil && existingAdmin.ID != adminId {
		return errors.New("this username is already used by another admin")
	}
	return nil
}

func validateUniqueEmailAndUsername(email, username string, repo repository.UserRepository) error {
	if err := validateUniqueEmail(email, 0, repo); err != nil {
		return err
	}
	return validateUniqueUsername(username, 0, repo)
}
