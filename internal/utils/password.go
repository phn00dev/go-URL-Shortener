package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword)
}
