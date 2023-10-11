package repository

import (
	"context"

	"github.com/catness812/PAD-lab1/journal_svc/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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
	newEntry := bson.M{
		"username": entry.Username,
		"title":    entry.Title,
		"content":  entry.Content,
	}
	_, err := repo.db.InsertOne(context.TODO(), newEntry)
	if err != nil {
		return err
	}
	return err
}
