package domain

import (
	"github.com/stretchr/testify/mock"
)

type MockAuthorRepository struct {
	mock.Mock
}

func NewMockAuthorRepository(mock mock.Mock) *MockAuthorRepository {
	return &MockAuthorRepository{
		Mock: mock,
	}
}

func (m *MockAuthorRepository) GetAll(limit int, offset int) ([]*Author, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]*Author), args.Error(1)
}

func (m *MockAuthorRepository) GetByName(name string) (*Author, error) {
	args := m.Called(name)
	return args.Get(0).(*Author), args.Error(1)
}

func (m *MockAuthorRepository) GetByID(ID int) (*Author, error) {
	args := m.Called(ID)
	return args.Get(0).(*Author), args.Error(1)
}

func (m *MockAuthorRepository) Create(author *Author) (*Author, error) {
	args := m.Called(author)
	return args.Get(0).(*Author), args.Error(1)
}

func (m *MockAuthorRepository) Update(ID int, author *Author) (*Author, error) {
	args := m.Called(ID, author)
	return args.Get(0).(*Author), args.Error(1)
}

func (m *MockAuthorRepository) Delete(ID int) error {
	args := m.Called(ID)
	return args.Error(0)
}
