package service

import "github.com/catness812/PAD-labs/journal_svc/internal/models"

type IJournalRepository interface {
	Save(entry *models.JournalEntry) error
	GetEntries(username string) ([]models.JournalEntry, error)
	DeleteEntries(username string) error
}

type JournalService struct {
	repo IJournalRepository
}

func InitJournalService(repo IJournalRepository) *JournalService {
	return &JournalService{
		repo: repo,
	}
}

func (svc *JournalService) RegisterNewEntry(entry models.JournalEntry) error {
	return svc.repo.Save(&entry)
}

func (svc *JournalService) GetUserEntries(username string) ([]models.JournalEntry, error) {
	return svc.repo.GetEntries(username)
}

func (svc *JournalService) DeleteUserEntries(username string) error {
	return svc.repo.DeleteEntries(username)
}
