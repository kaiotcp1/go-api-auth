package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI            string
	MongoDBName         string
	MongoUserCollection string
	HTTPPort            string
	JWTSecret           string
	JWTIssuer           string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{
		MongoURI:            os.Getenv("MONGO_URI"),
		MongoDBName:         os.Getenv("MONGO_DB_NAME"),
		MongoUserCollection: os.Getenv("MONGO_USER_COLLECTION"),
		HTTPPort:            os.Getenv("HTTP_PORT"),
		JWTSecret:           os.Getenv("JWT_SECRET"),
		JWTIssuer:           os.Getenv("JWT_ISSUER"),
	}

	if config.HTTPPort == "" {
		config.HTTPPort = "8080"
	}

	if config.MongoDBName == "" {
		return nil, fmt.Errorf("MONGO_DB_NAME is required")
	}

	if config.MongoUserCollection == "" {
		return nil, fmt.Errorf("MONGO_USER_COLLECTION is required")
	}

	return config, nil

}
