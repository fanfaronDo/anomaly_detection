package repository

import (
	entry "main/pkg/api/api/proto"

	"gorm.io/gorm"
)

type AnomalyItem interface {
	Create(entry.DataEntry)
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
