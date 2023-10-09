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

func (s *Server) RegisterNewUser(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newUser := models.User{
		Username: req.User.Username,
		Password: req.User.Password,
	}

	if err := s.UserService.RegisterUser(newUser); err != nil {
		slog.Errorf("Error registering new user: %v", err)
		return nil, err
	}

	slog.Info("User \"%v\" successfully created", newUser.Username)
	return &pb.RegisterResponse{
		Message: fmt.Sprintf("User \"%v\" successfully signed up", newUser.Username),
	}, nil
}
