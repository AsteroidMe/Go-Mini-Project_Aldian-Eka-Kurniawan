package repository

import (
	"eco-journal/entities"

	"gorm.io/gorm"
)

type AuthorRepoInterface interface {
	Create(author *entities.Author) (*entities.Author, error)
	Update(author *entities.Author) (*entities.Author, error)
	Delete(id uint) error
	FindAll(authors *[]entities.Author, page, limit int) error
	FindByID(id uint) (*entities.Author, error)
	Count() (int64, error)
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) Create(author *entities.Author) (*entities.Author, error) {
	if err := r.db.Create(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (r *authorRepository) Update(author *entities.Author) (*entities.Author, error) {
	if err := r.db.Save(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (r *authorRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Author{}, id).Error
}

func (r *authorRepository) FindAll(authors *[]entities.Author, page, limit int) error {
	offset := (page - 1) * limit
	return r.db.Offset(offset).Limit(limit).Find(authors).Error
}

func (r *authorRepository) FindByID(id uint) (*entities.Author, error) {
	var author entities.Author
	if err := r.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *authorRepository) Count() (int64, error) {
	var total int64
	err := r.db.Model(&entities.Author{}).Count(&total).Error
	return total, err
}
