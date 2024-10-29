package config

import (
	"os"
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
