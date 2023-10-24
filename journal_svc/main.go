package main

import (
	"fmt"
	"net"

	"github.com/catness812/PAD-lab1/journal_svc/internal/config"
	rpctransport "github.com/catness812/PAD-lab1/journal_svc/internal/controller/rpc-transport"
	"github.com/catness812/PAD-lab1/journal_svc/internal/pb"
	"github.com/catness812/PAD-lab1/journal_svc/internal/repository"
	"github.com/catness812/PAD-lab1/journal_svc/internal/service"
	"github.com/catness812/PAD-lab1/journal_svc/internal/utils"
	"github.com/catness812/PAD-lab1/journal_svc/pkg/db/mongo"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()
	db := mongo.LoadDatabase()
	journalRepo := repository.InitJournalRepository(db)
	journalSvc := service.InitJournalService(journalRepo)
	go grpcStart(journalSvc)
	utils.RegisterService()

	select {}
}

func grpcStart(journalSvc rpctransport.IJournalService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
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

	if err := s.Serve(lis); err != nil {
		slog.Error(err)
		panic(err)
	}
}
