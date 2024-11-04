package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbConfig   dbConfig   `json:"db_config"`
	HttpConfig httpConfig `json:"http_config"`
}

type dbConfig struct {
	DbHost     string `json:"db_host" env:"DB_HOST"`
	DbPort     string `json:"db_port" env:"DB_PORT"`
	DbUser     string `json:"db_user" env:"DB_USER"`
	DbPassword string `json:"db_password" env:"DB_PASSWORD"`
	DbName     string `json:"db_name" env:"DB_NAME"`
	DbSslMode  string `json:"db_sll_mode" env:"DB_SSL_MODE"`
	DbTimeZone string `json:"db_time_zone" env:"DB_TIME_ZONE"`
}

type httpConfig struct {
	HttpHost  string `json:"http_host" env:"HTTP_SERVER"`
	HttpPort  string `json:"http_port" env:"HTTP_PORT"`
	AppName   string `json:"app_name" env:"APP_NAME"`
	AppHeader string `json:"app_header" env:"APP_HEADER"`
}

func GetConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("../.env", &cfg)
	if err != nil {
		return nil, err
	}

	cfg = Config{
		DbConfig: dbConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
			DbSslMode:  os.Getenv("DB_SSL_MODE"),
			DbTimeZone: os.Getenv("DB_TIME_ZONE"),
		},
		HttpConfig: httpConfig{
			HttpHost:  os.Getenv("HTTP_HOST"),
			HttpPort:  os.Getenv("HTTP_PORT"),
			AppName:   os.Getenv("APP_NAME"),
			AppHeader: os.Getenv("APP_HEADER"),
		},
	}
	return &cfg, err
}
