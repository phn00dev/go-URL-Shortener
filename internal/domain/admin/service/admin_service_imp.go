package service

import (
	"errors"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
	"github.com/phn00dev/go-URL-Shortener/internal/utils"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: repo,
	}
}

func (s adminServiceImp) FindOneById(adminId int) (*model.Admin, error) {
	return s.adminRepo.GetOneById(adminId)
}

func (s adminServiceImp) FindAll() ([]model.Admin, error) {
	return s.adminRepo.GetAll()
}

func (s adminServiceImp) Create(createRequest dto.CreateAdminRequest) error {
	if err := validateUniqueEmailAndUsername(createRequest.Email, createRequest.Username, s.adminRepo); err != nil {
		return err
	}

	admin := model.Admin{
		Username:     createRequest.Username,
		Email:        createRequest.Email,
		PasswordHash: utils.HashPassword(createRequest.Password),
	}

	return s.adminRepo.Create(admin)
}

func (s adminServiceImp) Update(adminId int, updateRequest dto.UpdateAdminRequest) error {
	admin, err := s.adminRepo.GetOneById(adminId)
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New("admin not found")
	}

	if updateRequest.Username != "" {
		if err := validateUniqueUsername(updateRequest.Username, adminId, s.adminRepo); err != nil {
			return err
		}
		admin.Username = updateRequest.Username
	}

	if updateRequest.Email != "" {
		if err := validateUniqueEmail(updateRequest.Email, adminId, s.adminRepo); err != nil {
			return err
		}
		admin.Email = updateRequest.Email
	}

	return s.adminRepo.Update(admin.ID, *admin)
}

func (s adminServiceImp) Delete(adminId int) error {
	admin, err := s.adminRepo.GetOneById(adminId)
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New("admin not found")
	}
	return s.adminRepo.Delete(admin.ID)
}

func validateUniqueEmail(email string, adminId int, repo repository.AdminRepository) error {
	existingAdmin, err := repo.GetAdminByEmail(email)
	if err != nil {
		return err
	}
	if existingAdmin != nil && existingAdmin.ID != adminId {
		return errors.New("this email is already used by another admin")
	}
	return nil
}

func validateUniqueUsername(username string, adminId int, repo repository.AdminRepository) error {
	existingAdmin, err := repo.GetAdminByUsername(username)
	if err != nil {
		return err
	}
	if existingAdmin != nil && existingAdmin.ID != adminId {
		return errors.New("this username is already used by another admin")
	}
	return nil
}

func validateUniqueEmailAndUsername(email, username string, repo repository.AdminRepository) error {
	if err := validateUniqueEmail(email, 0, repo); err != nil {
		return err
	}
	return validateUniqueUsername(username, 0, repo)
}
