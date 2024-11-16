package service

import (
	"eco-journal/entities"
	"eco-journal/repository"
)

type JournalServiceInterface interface {
	Create(journal *entities.Journal) (*entities.Journal, error)
	Update(journal *entities.Journal) (*entities.Journal, error)
	Delete(id uint) error
	GetAll(page int, limit int) ([]entities.Journal, error)
	FindByID(id uint) (*entities.Journal, error)
	GetAuthorByID(id uint) (*entities.Author, error)
	GetCategoryByID(id uint) (*entities.Category, error)
	Count() (int64, error)
}

type journalService struct {
	journalRepo repository.JournalRepoInterface
}

func NewJournalService(journalRepo repository.JournalRepoInterface) *journalService {
	return &journalService{journalRepo}
}

func (s *journalService) Create(journal *entities.Journal) (*entities.Journal, error) {
	return s.journalRepo.Create(journal)
}

func (s *journalService) Update(journal *entities.Journal) (*entities.Journal, error) {
	return s.journalRepo.Update(journal)
}

func (s *journalService) Delete(id uint) error {
	return s.journalRepo.Delete(id)
}

func (s *journalService) GetAll(page int, limit int) ([]entities.Journal, error) {
	return s.journalRepo.FindAll(page, limit)
}

func (s *journalService) FindByID(id uint) (*entities.Journal, error) {
	return s.journalRepo.FindByID(id)
}

func (s *journalService) GetAuthorByID(id uint) (*entities.Author, error) {
	return s.journalRepo.GetAuthorByID(id)
}

func (s *journalService) GetCategoryByID(id uint) (*entities.Category, error) {
	return s.journalRepo.GetCategoryByID(id)
}

func (s *journalService) Count() (int64, error) {
	return s.journalRepo.Count()
}
