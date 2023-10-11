package repository

import (
	"github.com/catness812/PAD-lab1/journal_svc/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type JournalRepository struct {
	db *mongo.Collection
}

func InitJournalRepository(db *mongo.Collection) *JournalRepository {
	return &JournalRepository{
		db: db,
	}
}

func (repo *JournalRepository) Save(entry *models.JournalEntry) error {

	return nil
}
