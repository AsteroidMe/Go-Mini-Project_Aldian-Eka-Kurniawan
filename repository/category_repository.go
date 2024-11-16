package repository

import (
	"eco-journal/entities"

	"gorm.io/gorm"
)

type CategoryRepoInterface interface {
	Create(category *entities.Category) (*entities.Category, error)
	Update(category *entities.Category) (*entities.Category, error)
	Delete(id uint) error
	FindAll(categories *[]entities.Category, page, limit int) error
	FindByID(id uint) (*entities.Category, error)
	Count() (int64, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *entities.Category) (*entities.Category, error) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Update(category *entities.Category) (*entities.Category, error) {
	if err := r.db.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Category{}, id).Error
}

func (r *categoryRepository) FindAll(categories *[]entities.Category, page, limit int) error {
	offset := (page - 1) * limit
	return r.db.Offset(offset).Limit(limit).Find(categories).Error
}

func (r *categoryRepository) FindByID(id uint) (*entities.Category, error) {
	var category entities.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Count() (int64, error) {
	var total int64
	err := r.db.Model(&entities.Category{}).Count(&total).Error
	return total, err
}
