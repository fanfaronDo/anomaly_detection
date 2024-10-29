package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost                string
	DBPort                string
	DBSchema              string
	DBUser                string
	DBPassword            string
	DBSSLMode             string
	ServerTransmitterHost string
	ServerTransmitterPort string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Ошибка загрузки файла .env: %v", err)
	}

	return &Config{
		DBHost:                os.Getenv("DB_HOST"),
		DBPort:                os.Getenv("DB_PORT"),
		DBSchema:              os.Getenv("DB_SCHEMA"),
		DBUser:                os.Getenv("DB_USER"),
		DBPassword:            os.Getenv("DB_PASSWORD"),
		DBSSLMode:             os.Getenv("DB_SSLMODE"),
		ServerTransmitterHost: os.Getenv("SERVER_TRANSMITTER_HOST"),
		ServerTransmitterPort: os.Getenv("SERVER_TRANSMITTER_PORT"),
	}, nil
}
