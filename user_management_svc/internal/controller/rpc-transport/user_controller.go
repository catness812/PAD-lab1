package rpctransport

import "github.com/catness812/PAD-lab1/user_management_svc/internal/pb"

type IUserService interface{}

type Server struct {
	pb.UserServiceServer
	UserService IUserService
}
