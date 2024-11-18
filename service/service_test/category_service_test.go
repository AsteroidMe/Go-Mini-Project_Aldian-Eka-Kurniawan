package service_test

import (
	"testing"

	"eco-journal/entities"
	"eco-journal/repository/mocks"
	"eco-journal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCategoryService_Create(t *testing.T) {
	mockRepo := new(mocks.CategoryRepo)
	categoryService := service.NewCategoryService(mockRepo)

	category := &entities.Category{Name: "Test Category", Description: "Test Description"}
	mockRepo.On("Create", category).Return(category, nil)

	result, err := categoryService.Create(category)

	assert.NoError(t, err)
	assert.Equal(t, category, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_Update(t *testing.T) {
	mockRepo := new(mocks.CategoryRepo)
	categoryService := service.NewCategoryService(mockRepo)

	category := &entities.Category{ID: 1, Name: "Updated Category", Description: "Updated Description"}
	mockRepo.On("Update", category).Return(category, nil)

	result, err := categoryService.Update(category)

	assert.NoError(t, err)
	assert.Equal(t, category, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_Delete(t *testing.T) {
	mockRepo := new(mocks.CategoryRepo)
	categoryService := service.NewCategoryService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := categoryService.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_GetAll(t *testing.T) {
	mockRepo := new(mocks.CategoryRepo)
	categoryService := service.NewCategoryService(mockRepo)

	categories := []entities.Category{
		{ID: 1, Name: "Category 1", Description: "Description 1"},
		{ID: 2, Name: "Category 2", Description: "Description 2"},
	}
	mockRepo.On("Count").Return(int64(len(categories)), nil)
	mockRepo.On("FindAll", mock.Anything, 1, 10).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.Category)
		*arg = categories
	}).Return(nil)

	result, pagination, err := categoryService.GetAll(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, categories, result)
	assert.Equal(t, 1, pagination.CurrentPage)
	assert.Equal(t, 1, pagination.TotalPages)
	assert.Equal(t, int64(len(categories)), pagination.TotalItems)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_FindByID(t *testing.T) {
	mockRepo := new(mocks.CategoryRepo)
	categoryService := service.NewCategoryService(mockRepo)

	category := &entities.Category{ID: 1, Name: "Category 1", Description: "Description 1"}
	mockRepo.On("FindByID", uint(1)).Return(category, nil)

	result, err := categoryService.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, category, result)
	mockRepo.AssertExpectations(t)
}
