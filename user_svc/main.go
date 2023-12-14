package main

import (
	"fmt"
	"net"
	"net/http"
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
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	srvMetrics := srvMetrics(port)
	opts = append(opts, grpc.ChainUnaryInterceptor(
		srvMetrics.UnaryServerInterceptor(),
	))
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

func srvMetrics(port int) *grpcprom.ServerMetrics {
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)

	http.Handle(fmt.Sprintf("/metrics-user-%d", port), promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port+1000), nil); err != nil {
			slog.Error(err)
			panic(err)
		}
	}()
	return srvMetrics
}
