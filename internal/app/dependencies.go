package app

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/pkg/config"
	dbconfig "github.com/phn00dev/go-URL-Shortener/pkg/database/db_config"
	"github.com/phn00dev/go-URL-Shortener/pkg/httpClient"
)

type Dependencies struct {
	DB         *gorm.DB
	HttpClient *http.Client
	Config     *config.Config
}

func GetDependencies() (*Dependencies, error) {
	// Konfigurasiýany almak
	getConfig, err := config.GetConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}

	// Maglumatlar bazasy üçin konfigurasiýa döretmek
	newDB := dbconfig.NewDbConfig(getConfig)
	getDB, err := newDB.GetDbConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	fmt.Println("database connection successfully")

	// db seeder

	// newSeeder := seeders.NewDBSeeder(getDB)
	// if err := newSeeder.GetAllSeeder(); err != nil {
	// 	fmt.Printf("seeder error: %v", err.Error())
	// 	return nil, err
	// }

	// HTTP müşderisini döretmek
	clientHttp := httpClient.NewHttp()

	return &Dependencies{
		DB:         getDB,
		HttpClient: clientHttp,
		Config:     getConfig,
	}, nil
}
