package dto

import "github.com/phn00dev/go-URL-Shortener/internal/model"

type AdminLoginResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AdminRole   string `json:"admin_role"`
	AccessToken string `json:"access_token"`
}

func NewAdminLoginResponse(admin *model.Admin, accessToken string) *AdminLoginResponse {
	return &AdminLoginResponse{
		ID:          admin.ID,
		Username:    admin.Username,
		Email:       admin.Email,
		AdminRole:   admin.AdminRole,
		AccessToken: accessToken,
	}
}
