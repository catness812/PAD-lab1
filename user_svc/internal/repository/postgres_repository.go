package repository

import (
	"github.com/catness812/PAD-lab1/user_management_svc/internal/models"
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

func (repo *UserRepository) Save(user *models.User) error {
	err := repo.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
