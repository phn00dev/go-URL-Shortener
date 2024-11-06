package dto

type RegisterUserRequest struct {
	Username        string `json:"username" binding:"required,min=3,max=50" validate:"required,min=3,max=50"`
	Email           string `json:"email" binding:"required,email" validate:"required,email"`
	Password        string `json:"password" binding:"required,min=6" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" validate:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email" validate:"required,email"`
}

type UpdateUserPassword struct {
	OldPassword     string `json:"old_password" binding:"required" validate:"required"`
	Password        string `json:"password" binding:"required,min=6" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6" validate:"required,min=6"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}
