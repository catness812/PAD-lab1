package repository

import (
	"github.com/catness812/PAD-labs/user_svc/internal/models"
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

func (repo *UserRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) DeleteUserByUsername(username string) error {
	err := repo.db.Unscoped().Where("username = ?", username).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
