package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/catness812/PAD-labs/journal_svc/internal/config"
	rpctransport "github.com/catness812/PAD-labs/journal_svc/internal/controller/rpc-transport"
	"github.com/catness812/PAD-labs/journal_svc/internal/pb"
	"github.com/catness812/PAD-labs/journal_svc/internal/repository"
	"github.com/catness812/PAD-labs/journal_svc/internal/service"
	"github.com/catness812/PAD-labs/journal_svc/internal/utils"
	"github.com/catness812/PAD-labs/journal_svc/pkg/db/mongo"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()
	db := mongo.LoadDatabase()
	journalRepo := repository.InitJournalRepository(db)
	journalSvc := service.InitJournalService(journalRepo)

	for port := 50055; port <= 50057; port++ {
		go grpcStart(journalSvc, port)
		utils.RegisterService(port)
	}

	select {}
}

func grpcStart(journalSvc rpctransport.IJournalService, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(20),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterJournalServiceServer(s, &rpctransport.Server{
		JournalService: journalSvc,
	})

	slog.Infof("gRPC Server listening at %v\n", lis.Addr())

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.StoreInt32(&rpctransport.JournalPingCounter, 0)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			load := atomic.LoadInt32(&rpctransport.JournalPingCounter)
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
