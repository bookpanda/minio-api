package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port   string
	Env    string
	ApiKey string
}

type StoreConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type CorsConfig struct {
	AllowOrigins string
}

type Config struct {
	App   AppConfig
	Store StoreConfig
	Cors  CorsConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	appConfig := AppConfig{
		Port:   os.Getenv("APP_PORT"),
		Env:    os.Getenv("APP_ENV"),
		ApiKey: os.Getenv("APP_API_KEY"),
	}

	storeConfig := StoreConfig{
		Endpoint:  os.Getenv("STORE_ENDPOINT"),
		AccessKey: os.Getenv("STORE_ACCESS_KEY"),
		SecretKey: os.Getenv("STORE_SECRET_KEY"),
		UseSSL:    os.Getenv("STORE_USE_SSL") == "true",
	}

	corsConfig := CorsConfig{
		AllowOrigins: os.Getenv("CORS_ORIGINS"),
	}

	return &Config{
		App:   appConfig,
		Store: storeConfig,
		Cors:  corsConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
