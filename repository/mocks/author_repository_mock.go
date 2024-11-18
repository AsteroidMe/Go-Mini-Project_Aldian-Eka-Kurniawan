package mocks

import (
	"eco-journal/entities"

	"github.com/stretchr/testify/mock"
)

type AuthorRepo struct {
	mock.Mock
}

func (m *AuthorRepo) Create(author *entities.Author) (*entities.Author, error) {
	args := m.Called(author)
	return args.Get(0).(*entities.Author), args.Error(1)
}

func (m *AuthorRepo) Update(author *entities.Author) (*entities.Author, error) {
	args := m.Called(author)
	return args.Get(0).(*entities.Author), args.Error(1)
}

func (m *AuthorRepo) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *AuthorRepo) FindByID(id uint) (*entities.Author, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Author), args.Error(1)
}

func (m *AuthorRepo) FindAll(authors *[]entities.Author, page, limit int) error {
	args := m.Called(authors, page, limit)
	return args.Error(0)
}

func (m *AuthorRepo) Count() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
