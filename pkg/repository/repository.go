package repository

import (
	entry "github.com/fanfaronDo/anomaly_detection/pkg/api/api/proto"

	"gorm.io/gorm"
)

type AnomalyItem interface {
	Create(entry *entry.DataEntry)
	GetAll() ([]entry.DataEntry, error)
}

type Repository struct {
	AnomalyItem
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AnomalyItem: NewAnomalyItemPostgres(db),
	}
}
