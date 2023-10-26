package postgres

import (
	"time"

	"github.com/catness812/PAD-lab1/user_svc/internal/models"
	"github.com/gookit/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() *gorm.DB {
	time.Sleep(3 * time.Second)
	db := connect()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		slog.Error(err)
	}

	return db
}

func connect() *gorm.DB {
	var err error

	dsn := "postgres://postgres:pass@journaling-app-postgres:5432/journaling-app-db"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error(err)
		panic(err)
	} else {
		slog.Info("Successfully connected to the Postgres database")
	}

	return database
}
