package postgres

import (
	"fmt"

	"github.com/catness812/PAD-lab1/user_management_svc/internal/config"
	"github.com/catness812/PAD-lab1/user_management_svc/internal/models"
	"github.com/gookit/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() *gorm.DB {
	db := connect()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		slog.Error(err)
	}

	return db
}

func connect() *gorm.DB {
	var err error

	dsn := fmt.Sprintf(`host=%s
	dbname=%s
	user=%s
	password=%s
	port=%d
	sslmode=disable`,
		config.Cfg.Database.Host,
		config.Cfg.Database.DBName,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Port,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error(err)
	} else {
		slog.Info("Successfully connected to the Postgres database")
	}

	return database
}
