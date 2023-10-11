package rpctransport

import (
	"context"
	"fmt"

	"github.com/catness812/PAD-lab1/user_svc/internal/models"
	"github.com/catness812/PAD-lab1/user_svc/internal/pb"
	"github.com/gookit/slog"
)

type IUserService interface {
	RegisterNewUser(user models.User) error
	FindUser(username string) (*models.User, error)
}

type Server struct {
	pb.UserServiceServer
	UserService IUserService
}

func (s *Server) RegisterUser(_ context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	newUser := models.User{
		Username: req.User.Username,
		Password: req.User.Password,
	}

	user, err := s.UserService.FindUser(req.User.Username)
	if err != nil {
		slog.Errorf("Error finding user: %v", err)
		return nil, err
	}

	if user.ID == 0 {
		if err := s.UserService.RegisterNewUser(newUser); err != nil {
			slog.Errorf("Error registering new user: %v", err)
			return nil, err
		}
	} else if user.ID != 0 {
		slog.Errorf("User '%v' has already signed up", user.Username)
		return &pb.RegisterUserResponse{
			Message: fmt.Sprintf("User '%v' has already signed up", user.Username),
		}, nil
	}

	slog.Infof("User '%v' successfully created", newUser.Username)
	return &pb.RegisterUserResponse{
		Message: fmt.Sprintf("User '%v' successfully signed up", newUser.Username),
	}, nil
}
