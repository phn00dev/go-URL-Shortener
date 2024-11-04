package dto

type CreateAdminRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password_hash"`
	ConfirmPassword string `json:"confirm_password"`
	AdminRole       string `json:"admin_role"`
}

type UpdateAminRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	AdminRole string `json:"admin_role"`
}

type ChangeAdminPassword struct {
	OldPassword     string `json:"old_password"`
	Password        string `json:"password_hash"`
	ConfirmPassword string `json:"confirm_password"`
}
