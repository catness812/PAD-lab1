package main

import (
	"fmt"
	"net"

	"github.com/catness812/PAD-lab1/user_management_svc/internal/config"
	rpctransport "github.com/catness812/PAD-lab1/user_management_svc/internal/controller/rpc-transport"
	"github.com/catness812/PAD-lab1/user_management_svc/internal/pb"
	"github.com/catness812/PAD-lab1/user_management_svc/internal/repository"
	"github.com/catness812/PAD-lab1/user_management_svc/internal/service"
	"github.com/catness812/PAD-lab1/user_management_svc/pkg/db/postgres"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()
	db := postgres.LoadDatabase()
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	grpcStart(userSvc)
}

func grpcStart(userSvc rpctransport.IUserService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &rpctransport.Server{
		UserService: userSvc,
	})

	slog.Infof("gRPC Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		slog.Error(err)
		panic(err)
	}
}
