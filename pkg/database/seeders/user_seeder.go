package seeders

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/model"
	"github.com/phn00dev/go-URL-Shortener/internal/utils"

)

func userSeeder(db *gorm.DB) error {
	var users []model.User
	password := utils.HashPassword("12345678")

	for i := 1; i <= 5; i++ {
		user := model.User{
			Username:     "admin" + strconv.Itoa(i),
			Email:        "admin" + strconv.Itoa(i) + "@gmail.com",
			PasswordHash: password,
		}
		users = append(users, user)
	}

	if err := db.Create(&users).Error; err != nil {
		return err
	}
	return nil
}
