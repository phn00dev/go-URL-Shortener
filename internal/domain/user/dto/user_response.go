package dto

import "github.com/phn00dev/go-URL-Shortener/internal/model"

type UserLoginResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func NewUserLoginResponse(user *model.User, accessToken string) *UserLoginResponse {
	return &UserLoginResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: accessToken,
	}
}
