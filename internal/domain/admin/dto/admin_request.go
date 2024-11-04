package dto

type CreateAdminRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" validate:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email" validate:"required,email"`
	Password string `json:"password" binding:"required,min=6" validate:"required,min=6"`
}

type UpdateAdminRequest struct {
	Username string `json:"username" binding:"min=3,max=50" validate:"omitempty,min=3,max=50"`
	Email    string `json:"email" binding:"email" validate:"omitempty,email"`
}

type ChangeAdminPassword struct {
	OldPassword     string `json:"old_password"`
	Password        string `json:"password_hash"`
	ConfirmPassword string `json:"confirm_password"`
}
