package service

import "github.com/catness812/PAD-lab1/user_management_svc/internal/models"

type IUserRepository interface {
	Save(user *models.User) error
}

type UserService struct {
	repo IUserRepository
}

func InitUserService(repo IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) RegisterUser(user models.User) error {
	return svc.repo.Save(&user)
}
