package db

import (
	"github.com/fanfaronDo/anomaly_detection/pkg/config"
	"github.com/fanfaronDo/anomaly_detection/pkg/repository"
	"gorm.io/gorm"
)

type Connector interface {
	Connect() (*gorm.DB, error)
}

type DB struct {
	Connector
}

func NewBD(conf *config.Config) *DB {
	return &DB{
		Connector: repository.NewPostgresConnector(conf),
	}
}
