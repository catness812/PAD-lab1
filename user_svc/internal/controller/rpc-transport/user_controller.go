package rpctransport

import (
	"context"
	"fmt"

	"github.com/catness812/PAD-lab1/user_management_svc/internal/models"
	"github.com/catness812/PAD-lab1/user_management_svc/internal/pb"
	"github.com/gookit/slog"
)

type IUserService interface {
	RegisterUser(user models.User) error
}

type Server struct {
	pb.UserServiceServer
	UserService IUserService
}

func (s *Server) RegisterUser(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newUser := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := s.UserService.RegisterUser(newUser); err.Error() == "user has already signed up" {
		slog.Errorf("User '%v' has already signed up", newUser.Username)
		return &pb.RegisterResponse{
			Message: fmt.Sprintf("User '%v' has already signed up", newUser.Username),
		}, nil
	} else if err != nil {
		slog.Errorf("Error registering new user: %v", err)
		return nil, err
	}

	slog.Infof("User '%v' successfully created", newUser.Username)
	return &pb.RegisterResponse{
		Message: fmt.Sprintf("User '%v' successfully signed up", newUser.Username),
	}, nil
}
