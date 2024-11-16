package repository

import (
	"eco-journal/entities"

	"gorm.io/gorm"
)

type JournalRepoInterface interface {
	Create(journal *entities.Journal) (*entities.Journal, error)
	Update(journal *entities.Journal) (*entities.Journal, error)
	Delete(id uint) error
	FindAll() ([]entities.Journal, error)
	FindByID(id uint) (*entities.Journal, error)
	GetAuthorByID(id uint) (*entities.Author, error)
	GetCategoryByID(id uint) (*entities.Category, error)
}

type journalRepository struct {
	db *gorm.DB
}

func NewJournalRepository(db *gorm.DB) *journalRepository {
	return &journalRepository{db}
}

func (r *journalRepository) Create(journal *entities.Journal) (*entities.Journal, error) {
	if err := r.db.Create(journal).Error; err != nil {
		return nil, err
	}
	return journal, nil
}

func (r *journalRepository) Update(journal *entities.Journal) (*entities.Journal, error) {
	if err := r.db.Save(journal).Error; err != nil {
		return nil, err
	}
	return journal, nil
}

func (r *journalRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Journal{}, id).Error
}

func (r *journalRepository) FindAll() ([]entities.Journal, error) {
	var journals []entities.Journal
	if err := r.db.Preload("Author").Preload("Category").Find(&journals).Error; err != nil {
		return nil, err
	}
	return journals, nil
}

func (r *journalRepository) FindByID(id uint) (*entities.Journal, error) {
	var journal entities.Journal
	if err := r.db.Preload("Author").Preload("Category").First(&journal, id).Error; err != nil {
		return nil, err
	}
	return &journal, nil
}

func (r *journalRepository) GetAuthorByID(id uint) (*entities.Author, error) {
	var author entities.Author
	if err := r.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *journalRepository) GetCategoryByID(id uint) (*entities.Category, error) {
	var category entities.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
