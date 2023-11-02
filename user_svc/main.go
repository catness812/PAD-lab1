package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/catness812/PAD-labs/user_svc/internal/config"
	rpctransport "github.com/catness812/PAD-labs/user_svc/internal/controller/rpc-transport"
	"github.com/catness812/PAD-labs/user_svc/internal/pb"
	"github.com/catness812/PAD-labs/user_svc/internal/repository"
	"github.com/catness812/PAD-labs/user_svc/internal/service"
	"github.com/catness812/PAD-labs/user_svc/internal/utils"
	"github.com/catness812/PAD-labs/user_svc/pkg/db/postgres"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()
	db := postgres.LoadDatabase()
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)

	for port := 50051; port <= 50053; port++ {
		go grpcStart(userSvc, port)
		utils.RegisterService(port)
	}

	select {}
}

func grpcStart(userSvc rpctransport.IUserService, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(20),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, &rpctransport.Server{
		UserService: userSvc,
	})

	slog.Infof("gRPC Server listening at %v\n", lis.Addr())

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.StoreInt32(&rpctransport.UserPingCounter, 0)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			load := atomic.LoadInt32(&rpctransport.UserPingCounter)
			if load > 60 {
				slog.Errorf("Critical load reached: %d pings per second", load)
			}
		}
	}()

	if err := s.Serve(lis); err != nil {
		slog.Error(err)
		panic(err)
	}
}
