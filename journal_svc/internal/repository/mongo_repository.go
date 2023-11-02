package repository

import (
	"context"

	"github.com/catness812/PAD-labs/journal_svc/internal/models"
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

func (repo *JournalRepository) GetEntries(username string) ([]models.JournalEntry, error) {
	filter := bson.M{"username": username}
	cursor, err := repo.db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var entries []models.JournalEntry
	for cursor.Next(context.TODO()) {
		var entry models.JournalEntry
		err := cursor.Decode(&entry)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
