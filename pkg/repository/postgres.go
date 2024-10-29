package repository

import (
	"fmt"
	"log"

	"github.com/fanfaronDo/anomaly_detection/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnector struct {
	conf *config.Config
}

func (p *PostgresConnector) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
		p.conf.DBHost,
		p.conf.DBUser,
		p.conf.DBSchema,
		p.conf.DBPassword,
		p.conf.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Not connected to database %v\n", err)
		return db, err
	}
	return db, nil
}

func NewPostgresConnector(config *config.Config) *PostgresConnector {
	return &PostgresConnector{config}
}
