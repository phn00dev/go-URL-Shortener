package dbconfig

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/pkg/config"
)

type DbConfig struct {
	config *config.Config
}

func NewDbConfig(config *config.Config) *DbConfig {
	return &DbConfig{
		config: config,
	}
}

func (dbConfig *DbConfig) GetDbConfig() (*gorm.DB, error) {
	// DSN düzmek
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ashgabat",
		dbConfig.config.DbConfig.DbHost,
		dbConfig.config.DbConfig.DbUser,
		dbConfig.config.DbConfig.DbPassword,
		dbConfig.config.DbConfig.DbName,
		dbConfig.config.DbConfig.DbPort,
	)

	// PostgreSQL bilen baglanyşmak
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("database connection error: %v", err)
		return nil, err
	}
	return db, nil
}
