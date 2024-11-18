package service_test

import (
	"testing"

	"eco-journal/entities"
	"eco-journal/repository/mocks"
	"eco-journal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthorService_Create(t *testing.T) {
	mockRepo := new(mocks.AuthorRepo)
	authorService := service.NewAuthorService(mockRepo)

	author := &entities.Author{FirstName: "John", LastName: "Doe", Bio: "Author Bio"}
	mockRepo.On("Create", author).Return(author, nil)

	result, err := authorService.Create(author)

	assert.NoError(t, err)
	assert.Equal(t, author, result)
	mockRepo.AssertExpectations(t)
}

func TestAuthorService_Update(t *testing.T) {
	mockRepo := new(mocks.AuthorRepo)
	authorService := service.NewAuthorService(mockRepo)

	author := &entities.Author{ID: 1, FirstName: "John", LastName: "Doe", Bio: "Updated Bio"}
	mockRepo.On("Update", author).Return(author, nil)

	result, err := authorService.Update(author)

	assert.NoError(t, err)
	assert.Equal(t, author, result)
	mockRepo.AssertExpectations(t)
}

func TestAuthorService_Delete(t *testing.T) {
	mockRepo := new(mocks.AuthorRepo)
	authorService := service.NewAuthorService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := authorService.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAuthorService_GetAll(t *testing.T) {
	mockRepo := new(mocks.AuthorRepo)
	authorService := service.NewAuthorService(mockRepo)

	authors := []entities.Author{
		{ID: 1, FirstName: "John", LastName: "Doe", Bio: "Author Bio"},
		{ID: 2, FirstName: "Jane", LastName: "Doe", Bio: "Author Bio"},
	}
	mockRepo.On("Count").Return(int64(len(authors)), nil)
	mockRepo.On("FindAll", mock.Anything, 1, 10).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.Author)
		*arg = authors
	}).Return(nil)

	result, pagination, err := authorService.GetAll(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, authors, result)
	assert.Equal(t, 1, pagination.CurrentPage)
	assert.Equal(t, 1, pagination.TotalPages)
	assert.Equal(t, int64(len(authors)), pagination.TotalItems)
	mockRepo.AssertExpectations(t)
}

func TestAuthorService_FindByID(t *testing.T) {
	mockRepo := new(mocks.AuthorRepo)
	authorService := service.NewAuthorService(mockRepo)

	author := &entities.Author{ID: 1, FirstName: "John", LastName: "Doe", Bio: "Author Bio"}
	mockRepo.On("FindByID", uint(1)).Return(author, nil)

	result, err := authorService.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, author, result)
	mockRepo.AssertExpectations(t)
}
