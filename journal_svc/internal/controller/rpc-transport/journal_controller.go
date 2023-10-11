package rpctransport

import (
	"context"
	"fmt"

	"github.com/catness812/PAD-lab1/journal_svc/internal/models"
	"github.com/catness812/PAD-lab1/journal_svc/internal/pb"
	"github.com/gookit/slog"
)

type IJournalService interface {
	RegisterNewEntry(entry models.JournalEntry) error
}

type Server struct {
	pb.JournalServiceServer
	JournalService IJournalService
}

func (s *Server) RegisterEntry(_ context.Context, req *pb.RegisterEntryRequest) (*pb.RegisterEntryResponse, error) {
	newEntry := models.JournalEntry{
		Username: req.Entry.Username,
		Title:    req.Entry.Title,
		Content:  req.Entry.Content,
	}

	if err := s.JournalService.RegisterNewEntry(newEntry); err != nil {
		slog.Errorf("Error registering new journal entry: %v", err)
		return nil, err
	}

	slog.Infof("Journal entry named '%v' for user '%v' successfully created", newEntry.Title, newEntry.Username)
	return &pb.RegisterEntryResponse{
		Message: fmt.Sprintf("Journal entry named '%v' for user '%v' successfully created", newEntry.Title, newEntry.Username),
	}, nil
}
