package service

import (
	"errors"
	"fmt"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
	"github.com/phn00dev/go-URL-Shortener/internal/utils"
	jwttoken "github.com/phn00dev/go-URL-Shortener/pkg/jwtToken"
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
	existingUser, err := s.userRepo.FindByUsernameOrEmail(createRequest.Username, createRequest.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("username or email already exists")
	}

	newUser := model.User{
		Username:     createRequest.Username,
		Email:        createRequest.Email,
		PasswordHash: utils.HashPassword(createRequest.Password),
	}
	return s.userRepo.Create(newUser)
}

func (s userServiceImp) Update(userId int, updateRequest dto.UpdateUserRequest) error {
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		return err
	}
	existingUserEmail, err := s.userRepo.GetByEmail(updateRequest.Email)
	if err == nil && existingUserEmail.ID != userId {
		// Eger email başga admin tarapyndan eýýelenýän bolsa, ýalňyşlyk döretmek
		return fmt.Errorf("e-mail salgy eýýäm ulanylýar: %s", updateRequest.Email)
	}

	existingAdminUsername, err := s.userRepo.GetByUsername(updateRequest.Username)
	if err == nil && existingAdminUsername.ID != userId {
		// Eger username başga admin tarapyndan eýýelenýän bolsa, ýalňyşlyk döretmek
		return fmt.Errorf("username ady eýýäm ulanylýar: %s", updateRequest.Username)
	}
	user.Username = updateRequest.Username
	user.Email = updateRequest.Email
	return s.userRepo.Update(user.ID, *user)
}

func (s userServiceImp) Delete(userId int) error {
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("something error")
	}
	return s.userRepo.Delete(user.ID)
}

func (s userServiceImp) UpdateUserPassword(userId int, updatePasswordRequest dto.UpdateUserPassword) error {
	user, err := s.userRepo.GetById(userId)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("user not found")
	}
	if !utils.CheckPasswordHash(updatePasswordRequest.OldPassword, user.PasswordHash) {
		return errors.New("old password is incorrect")
	}
	if updatePasswordRequest.Password != updatePasswordRequest.ConfirmPassword {
		return errors.New("new password and confirmation do not match")
	}
	newPasswordHash := utils.HashPassword(updatePasswordRequest.Password)
	return s.userRepo.UpdateUserPassword(user.ID, newPasswordHash)

}

func (s userServiceImp) LoginUser(loginRequest dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	// get user with email
	user, err := s.userRepo.GetByUsername(loginRequest.Username)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("username or password wrong")
	}
	// password barlag
	if !utils.CheckPasswordHash(loginRequest.Password, user.PasswordHash) {
		return nil, errors.New("username or password wrong")
	}
	// generate token
	accessToken, err := jwttoken.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, errors.New("something wrong")
	}
	loginResponse := dto.NewUserLoginResponse(user, accessToken)
	return loginResponse, nil
}
