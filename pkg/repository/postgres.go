package repository

import (
	"fmt"
	"log"
	"main/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnector(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DBHost,
		config.DBUser,
		config.DBSchema,
		config.DBPassword,
		config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Not connected to database %v\n", err)
		return db, err
	}
	return db, nil
}
