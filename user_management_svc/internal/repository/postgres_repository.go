package repository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
