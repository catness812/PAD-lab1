package rpctransport

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/catness812/PAD-labs/journal_svc/internal/models"
	"github.com/catness812/PAD-labs/journal_svc/internal/pb"
	usersvc "github.com/catness812/PAD-labs/journal_svc/internal/user_svc"
	"github.com/gookit/slog"
)

var JournalPingCounter int32

type IJournalService interface {
	RegisterNewEntry(entry models.JournalEntry) error
	GetUserEntries(username string) ([]models.JournalEntry, error)
}

type Server struct {
	pb.JournalServiceServer
	JournalService IJournalService
}

func (s *Server) RegisterEntry(_ context.Context, req *pb.RegisterEntryRequest) (*pb.RegisterEntryResponse, error) {
	atomic.AddInt32(&JournalPingCounter, 1)
	if _, err := usersvc.UserServiceClient().CheckIfUserExists(context.Background(), &pb.User{Username: req.Entry.Username}); err != nil {
		slog.Errorf("User '%v' not found", req.Entry.Username)
		return &pb.RegisterEntryResponse{
			Message: fmt.Sprintf("User '%v' not found", req.Entry.Username),
		}, nil
	}

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

func (s *Server) GetUserEntries(_ context.Context, req *pb.GetUserEntriesRequest) (*pb.GetUserEntriesResponse, error) {
	atomic.AddInt32(&JournalPingCounter, 1)
	if _, err := usersvc.UserServiceClient().CheckIfUserExists(context.Background(), &pb.User{Username: req.Username}); err != nil {
		slog.Errorf("User '%v' not found", req.Username)
		return &pb.GetUserEntriesResponse{
			Message: fmt.Sprintf("User '%v' not found", req.Username),
		}, nil
	}

	entries, err := s.JournalService.GetUserEntries(req.Username)

	if err != nil {
		slog.Errorf("Error retrieving journal entries for user '%v': %v", req.Username, err)
		return &pb.GetUserEntriesResponse{
			Message: fmt.Sprintf("Error retrieving journal entries for user '%v': %v", req.Username, err),
		}, nil
	}

	var pbEntries []*pb.Entry
	for _, entry := range entries {
		pbEntry := &pb.Entry{
			Username: entry.Username,
			Title:    entry.Title,
			Content:  entry.Content,
		}
		pbEntries = append(pbEntries, pbEntry)
	}

	slog.Infof("Successfully retrieved journal entries for user '%v'", req.Username)
	return &pb.GetUserEntriesResponse{
		Message: fmt.Sprintf("Successfully retrieved journal entries for user '%v'", req.Username),
		Entries: pbEntries,
	}, nil
}
