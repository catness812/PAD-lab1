package rpctransport

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/catness812/PAD-lab1/user_svc/internal/models"
	"github.com/catness812/PAD-lab1/user_svc/internal/pb"
	"github.com/catness812/PAD-lab1/user_svc/internal/utils"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gookit/slog"
)

var UserPingCounter int32

type IUserService interface {
	RegisterNewUser(user models.User) error
	FindUser(username string) (*models.User, error)
	DeleteUser(username string) error
}

type Server struct {
	pb.UserServiceServer
	UserService IUserService
}

func (s *Server) RegisterUser(_ context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	atomic.AddInt32(&UserPingCounter, 1)
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
		slog.Infof("User '%v' successfully created", newUser.Username)
		return &pb.RegisterUserResponse{
			Message: fmt.Sprintf("User '%v' successfully signed up", newUser.Username),
		}, nil
	}

	if user.ID != 0 {
		slog.Errorf("User '%v' has already signed up", user.Username)
		return &pb.RegisterUserResponse{
			Message: fmt.Sprintf("User '%v' has already signed up", user.Username),
		}, nil
	}

	return nil, nil
}

func (s *Server) CheckIfUserExists(_ context.Context, req *pb.User) (*empty.Empty, error) {
	atomic.AddInt32(&UserPingCounter, 1)
	user, err := s.UserService.FindUser(req.Username)
	if err != nil {
		slog.Errorf("Error finding user: %v", err)
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &empty.Empty{}, nil
}

func (s *Server) DeleteUser(_ context.Context, req *pb.User) (*pb.DeleteUserResponse, error) {
	atomic.AddInt32(&UserPingCounter, 1)
	user, err := s.UserService.FindUser(req.Username)
	if err != nil {
		slog.Errorf("Error finding user: %v", err)
		return nil, err
	}

	if user.ID == 0 {
		slog.Errorf("User '%v' not found", req.Username)
		return &pb.DeleteUserResponse{
			Message: fmt.Sprintf("User '%v' not found", req.Username),
		}, nil
	}

	if err = utils.ValidatePassword(user.Password, req.Password); err != nil {
		slog.Errorf("Wrong password for user '%v'", req.Username)
		return &pb.DeleteUserResponse{
			Message: "Wrong password",
		}, nil
	}

	if err = s.UserService.DeleteUser(req.Username); err != nil {
		slog.Errorf("Error deleting user '%v': %v", req.Username, err)
		return nil, err
	}

	slog.Infof("User '%v' successfully deleted", req.Username)
	return &pb.DeleteUserResponse{
		Message: fmt.Sprintf("User '%v' successfully deleted", req.Username),
	}, nil
}
