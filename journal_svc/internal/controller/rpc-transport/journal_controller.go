package rpctransport

import (
	"context"

	"github.com/catness812/PAD-lab1/journal_svc/internal/models"
	"github.com/catness812/PAD-lab1/journal_svc/internal/pb"
)

type IJournalService interface {
	RegisterNewEntry(entry models.JournalEntry) error
}

type Server struct {
	pb.JournalServiceServer
	JournalService IJournalService
}

func (s *Server) RegisterEntry(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, nil
}
