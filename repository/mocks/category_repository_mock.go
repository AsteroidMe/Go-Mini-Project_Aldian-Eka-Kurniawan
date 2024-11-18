package mocks

import (
	"eco-journal/entities"

	"github.com/stretchr/testify/mock"
)

type CategoryRepo struct {
	mock.Mock
}

func (m *CategoryRepo) Create(category *entities.Category) (*entities.Category, error) {
	args := m.Called(category)
	return args.Get(0).(*entities.Category), args.Error(1)
}

func (m *CategoryRepo) Update(category *entities.Category) (*entities.Category, error) {
	args := m.Called(category)
	return args.Get(0).(*entities.Category), args.Error(1)
}

func (m *CategoryRepo) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *CategoryRepo) FindByID(id uint) (*entities.Category, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Category), args.Error(1)
}

func (m *CategoryRepo) FindAll(categories *[]entities.Category, page, limit int) error {
	args := m.Called(categories, page, limit)
	return args.Error(0)
}

func (m *CategoryRepo) Count() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
