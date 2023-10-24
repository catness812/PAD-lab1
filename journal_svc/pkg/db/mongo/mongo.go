package mongo

import (
	"context"
	"fmt"

	"github.com/catness812/PAD-lab1/journal_svc/internal/config"
	"github.com/gookit/slog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func LoadDatabase() *mongo.Collection {
	client := connect()
	usersCollection := client.Database(config.Cfg.Database.DBName).Collection("journal-entries")

	return usersCollection
}

func connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d",
			config.Cfg.Database.User,
			config.Cfg.Database.Password,
			config.Cfg.Host,
			config.Cfg.Database.Port)))

	if err != nil {
		slog.Error(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		slog.Error(err)
	} else {
		slog.Info("Successfully connected to the MongoDB database")
	}

	return client
}
