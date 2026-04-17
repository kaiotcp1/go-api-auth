package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName             string
	AppVersion          string
	MongoURI            string
	MongoDBName         string
	MongoUserCollection string
	HTTPPort            string
	JWTSecret           string
	JWTIssuer           string
	AllowedOrigins      string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{
		AppName:             os.Getenv("APP_NAME"),
		AppVersion:          os.Getenv("APP_VERSION"),
		MongoURI:            os.Getenv("MONGO_URI"),
		MongoDBName:         os.Getenv("MONGO_DB_NAME"),
		MongoUserCollection: os.Getenv("MONGO_USER_COLLECTION"),
		HTTPPort:            os.Getenv("HTTP_PORT"),
		JWTSecret:           os.Getenv("JWT_SECRET"),
		JWTIssuer:           os.Getenv("JWT_ISSUER"),
		AllowedOrigins:      os.Getenv("ALLOWED_ORIGINS"),
	}

	if config.AppName == "" {
		config.AppName = "Go Auth API"
	}

	if config.AppVersion == "" {
		config.AppVersion = "1.1.0"
	}

	if config.HTTPPort == "" {
		config.HTTPPort = "8080"
	}

	if config.AllowedOrigins == "" {
		config.AllowedOrigins = "*"
	}

	if config.MongoDBName == "" {
		return nil, fmt.Errorf("MONGO_DB_NAME is required")
	}

	if config.MongoUserCollection == "" {
		return nil, fmt.Errorf("MONGO_USER_COLLECTION is required")
	}

	return config, nil

}
