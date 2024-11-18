package service

import (
	"eco-journal/entities"
	"eco-journal/repository"
)

type AuthorServiceInterface interface {
	Create(author *entities.Author) (*entities.Author, error)
	Update(author *entities.Author) (*entities.Author, error)
	Delete(id uint) error
	GetAll(page, limit int) ([]entities.Author, *entities.Pagination, error)
	FindByID(id uint) (*entities.Author, error)
}

type authorService struct {
	authorRepo repository.AuthorRepoInterface
}

func NewAuthorService(authorRepo repository.AuthorRepoInterface) *authorService {
	return &authorService{authorRepo}
}

func (s *authorService) Create(author *entities.Author) (*entities.Author, error) {
	return s.authorRepo.Create(author)
}

func (s *authorService) Update(author *entities.Author) (*entities.Author, error) {
	return s.authorRepo.Update(author)
}

func (s *authorService) Delete(id uint) error {
	return s.authorRepo.Delete(id)
}

func (s *authorService) GetAll(page, limit int) ([]entities.Author, *entities.Pagination, error) {
	var authors []entities.Author

	totalItems, err := s.authorRepo.Count()
	if err != nil {
		return nil, nil, err
	}

	if err := s.authorRepo.FindAll(&authors, page, limit); err != nil {
		return nil, nil, err
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	pagination := &entities.Pagination{
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	return authors, pagination, nil
}

func (s *authorService) FindByID(id uint) (*entities.Author, error) {
	return s.authorRepo.FindByID(id)
}
