package repository

import (
	entry "main/pkg/api/api/proto"

	"gorm.io/gorm"
)

type AnomalyItemPostgres struct {
	db *gorm.DB
}

type Transmitter struct {
	SessionID string
	Frequence float64
	Timestamp int
}

func NewAnomalyItemPostgres(db *gorm.DB) *AnomalyItemPostgres {
	return &AnomalyItemPostgres{db: db}
}

func (item *AnomalyItemPostgres) Create(entry entry.DataEntry) {
	item.db.Create(&Transmitter{
		SessionID: entry.SessionId,
		Frequence: entry.Frequency,
		Timestamp: int(entry.Timestamp),
	})
}

func (item *AnomalyItemPostgres) GetAll() ([]entry.DataEntry, error) {
	var entrys []entry.DataEntry
	results := item.db.Find(&entrys)
	if results.Error != nil {
		return []entry.DataEntry{}, results.Error
	}

	return entrys, nil
}
