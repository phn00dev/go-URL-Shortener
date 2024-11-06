package service

import (
	"errors"
	"fmt"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
	"github.com/phn00dev/go-URL-Shortener/internal/utils"
	jwttoken "github.com/phn00dev/go-URL-Shortener/pkg/jwtToken"

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
	// Gaýtalanýan ulanyjy adyny ýa-da emaili barlamak
	existingAdmin, err := s.adminRepo.FindByUsernameOrEmail(createRequest.Username, createRequest.Email)
	if err != nil {
		return err
	}
	if existingAdmin != nil {
		return errors.New("username or email already exists")
	}

	// Täze admin döretmek
	admin := model.Admin{
		Username:     createRequest.Username,
		Email:        createRequest.Email,
		AdminRole:    createRequest.AdminRole,
		PasswordHash: utils.HashPassword(createRequest.Password),
	}
	return s.adminRepo.Create(admin)
}

func (s adminServiceImp) Update(adminId int, updateRequest dto.UpdateAdminRequest) error {
	admin, err := s.adminRepo.GetOneById(adminId)
	if err != nil {
		return err
	}
	existingAdminEmail, err := s.adminRepo.GetAdminByEmail(updateRequest.Email)
	if err == nil && existingAdminEmail.ID != adminId {
		// Eger email başga admin tarapyndan eýýelenýän bolsa, ýalňyşlyk döretmek
		return fmt.Errorf("e-mail salgy eýýäm ulanylýar: %s", updateRequest.Email)
	}

	existingAdminUsername, err := s.adminRepo.GetAdminByUsername(updateRequest.Username)
	if err == nil && existingAdminUsername.ID != adminId {
		// Eger username başga admin tarapyndan eýýelenýän bolsa, ýalňyşlyk döretmek
		return fmt.Errorf("username ady eýýäm ulanylýar: %s", updateRequest.Username)
	}

	admin.Username = updateRequest.Username
	admin.Email = updateRequest.Email
	admin.AdminRole = updateRequest.AdminRole
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

func (s adminServiceImp) UpdateAdminPassword(adminId int, changePasswordRequest dto.ChangeAdminPassword) error {

	admin, err := s.adminRepo.GetOneById(adminId)
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New("admin not found")
	}

	if !utils.CheckPasswordHash(changePasswordRequest.OldPassword, admin.PasswordHash) {
		return errors.New("old password is incorrect")
	}
	if changePasswordRequest.Password != changePasswordRequest.ConfirmPassword {
		return errors.New("new password and confirmation do not match")
	}
	newPasswordHash := utils.HashPassword(changePasswordRequest.Password)
	return s.adminRepo.UpdateAdminPassword(admin.ID, newPasswordHash)
}

func (s adminServiceImp) AdminLogin(loginRequest dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
	// get admin with username
	admin, err := s.adminRepo.GetAdminByUsername(loginRequest.Username)
	if err != nil {
		return nil, err
	}
	// check password
	if !utils.CheckPasswordHash(loginRequest.Password, admin.PasswordHash) {
		return nil, errors.New("username or password wrong")
	}
	// generate token

	accessToken, err := jwttoken.GenerateAdminToken(admin.ID, admin.Username, admin.Email, admin.AdminRole)
	if err != nil {
		return nil, err
	}
	loginResponse := dto.NewAdminLoginResponse(admin, accessToken)
	return loginResponse, nil
}
