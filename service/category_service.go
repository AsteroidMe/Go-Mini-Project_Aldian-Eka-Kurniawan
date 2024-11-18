package service

import (
	"eco-journal/entities"
	"eco-journal/repository"
)

type CategoryServiceInterface interface {
	Create(category *entities.Category) (*entities.Category, error)
	Update(category *entities.Category) (*entities.Category, error)
	Delete(id uint) error
	GetAll(page, limit int) ([]entities.Category, *entities.Pagination, error)
	FindByID(id uint) (*entities.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepoInterface
}

func NewCategoryService(categoryRepo repository.CategoryRepoInterface) *categoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) Create(category *entities.Category) (*entities.Category, error) {
	return s.categoryRepo.Create(category)
}

func (s *categoryService) Update(category *entities.Category) (*entities.Category, error) {
	return s.categoryRepo.Update(category)
}

func (s *categoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
}

func (s *categoryService) GetAll(page, limit int) ([]entities.Category, *entities.Pagination, error) {
	var categories []entities.Category

	totalItems, err := s.categoryRepo.Count()
	if err != nil {
		return nil, nil, err
	}

	if err := s.categoryRepo.FindAll(&categories, page, limit); err != nil {
		return nil, nil, err
	}

	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	pagination := &entities.Pagination{
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	return categories, pagination, nil
}

func (s *categoryService) FindByID(id uint) (*entities.Category, error) {
	return s.categoryRepo.FindByID(id)
}
