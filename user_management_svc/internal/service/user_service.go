package service

type IUserRepository interface{}

type UserService struct {
	repo IUserRepository
}

func InitUserService(repo IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
