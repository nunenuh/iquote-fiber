package domain

import (
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func NewMockCategoryRepository() *MockCategoryRepository {
	return &MockCategoryRepository{}
}

func (m *MockCategoryRepository) GetAll(limit int, offset int) ([]*Category, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]*Category), args.Error(1)
}

func (m *MockCategoryRepository) GetByName(name string) ([]*Category, error) {
	args := m.Called(name)
	return args.Get(0).([]*Category), args.Error(1)
}

func (m *MockCategoryRepository) GetByParentID(ID int) ([]*Category, error) {
	args := m.Called(ID)
	return args.Get(0).([]*Category), args.Error(1)
}

func (m *MockCategoryRepository) GetByID(ID int) (*Category, error) {
	args := m.Called(ID)
	return args.Get(0).(*Category), args.Error(1)
}

func (m *MockCategoryRepository) Create(category *Category) (*Category, error) {
	args := m.Called(category)
	return args.Get(0).(*Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(ID int, category *Category) (*Category, error) {
	args := m.Called(ID, category)
	return args.Get(0).(*Category), args.Error(1)
}

func (m *MockCategoryRepository) Delete(ID int) error {
	args := m.Called(ID)
	return args.Error(0)
}
