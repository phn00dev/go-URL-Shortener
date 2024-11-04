package model

import "time"

type Admin struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	AdminRole string    `json:"admin_role"` // super_admin, admin
	CreatedAt time.Time `json:"created_at"`
}
