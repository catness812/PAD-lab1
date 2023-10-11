package repository

import (
	"errors"

	"github.com/catness812/PAD-lab1/user_svc/internal/models"
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
	err := repo.db.Where("username=?", user.Username).Find(&user).Error
	if err != nil {
		return err
	}
	if user.ID == 0 {
		err := repo.db.Create(user).Error
		if err != nil {
			return err
		}
	} else {
		return errors.New("user has already signed up")
	}
	return nil
}
