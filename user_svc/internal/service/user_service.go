package service

import (
	"github.com/catness812/PAD-labs/user_svc/internal/models"
	"github.com/catness812/PAD-labs/user_svc/internal/utils"
)

type IUserRepository interface {
	Save(user *models.User) error
	FindUserByUsername(username string) (*models.User, error)
	DeleteUserByUsername(username string) error
}

type UserService struct {
	repo IUserRepository
}

func InitUserService(repo IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) RegisterNewUser(user models.User) error {
	user.Username = utils.CleanUsername(user.Username)
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return svc.repo.Save(&user)
}

func (svc *UserService) FindUser(username string) (*models.User, error) {
	return svc.repo.FindUserByUsername(username)
}

func (svc *UserService) DeleteUser(username string) error {
	return svc.repo.DeleteUserByUsername(username)
}
