package service

import "github.com/catness812/PAD-lab1/journal_svc/internal/models"

type IJournalRepository interface {
	Save(entry *models.JournalEntry) error
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
