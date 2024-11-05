package seeders

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/model"
	"github.com/phn00dev/go-URL-Shortener/internal/utils"
)

func adminSeeder(db *gorm.DB) error {
	var admins []model.Admin
	password := utils.HashPassword("12345678")

	superAdmin := model.Admin{
		Username:     "polat",
		Email:        "hudayberdipolatgmail.com",
		AdminRole:    "super_admin",
		PasswordHash: password,
	}

	admins = append(admins, superAdmin)
	for i := 1; i <= 5; i++ {
		admin := model.Admin{
			Username:     "admin" + strconv.Itoa(i),
			Email:        "admin" + strconv.Itoa(i) + "@gmail.com",
			AdminRole:    "admin",
			PasswordHash: password,
		}
		admins = append(admins, admin)
	}

	if err := db.Create(&admins).Error; err != nil {
		return err
	}
	return nil
}
