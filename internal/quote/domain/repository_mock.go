package domain

import (
	"github.com/stretchr/testify/mock"
)

type MockQuoteRepository struct {
	mock.Mock
}

func NewMockQuoteRepository() *MockQuoteRepository {
	return &MockQuoteRepository{}
}

func (m *MockQuoteRepository) GetAll(limit int, offset int) ([]*Quote, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]*Quote), args.Error(1)
}

func (m *MockQuoteRepository) GetByAuthorName(name string, limit int, offset int) ([]*Quote, error) {
	args := m.Called(name, limit, offset)
	return args.Get(0).([]*Quote), args.Error(1)
}

func (m *MockQuoteRepository) GetByAuthorID(ID int, limit int, offset int) ([]*Quote, error) {
	args := m.Called(ID, limit, offset)
	return args.Get(0).([]*Quote), args.Error(1)
}

func (m *MockQuoteRepository) GetByCategoryName(name string, limit int, offset int) ([]*Quote, error) {
	args := m.Called(name, limit, offset)
	return args.Get(0).([]*Quote), args.Error(1)
}

func (m *MockQuoteRepository) GetByCategoryID(ID int, limit int, offset int) ([]*Quote, error) {
	args := m.Called(ID, limit, offset)
	return args.Get(0).([]*Quote), args.Error(1)
}

func (m *MockQuoteRepository) Like(quoteID int, userID int) (*Quote, error) {
	args := m.Called(quoteID, userID)
	return args.Get(0).(*Quote), args.Error(1)
}

func (m *MockQuoteRepository) Unlike(quoteID int, userID int) (*Quote, error) {
	args := m.Called(quoteID, userID)
	return args.Get(0).(*Quote), args.Error(1)
}

func (m *MockQuoteRepository) GetByID(ID int) (*Quote, error) {
	args := m.Called(ID)
	return args.Get(0).(*Quote), args.Error(1)
}

func (m *MockQuoteRepository) Create(category *Quote) (*Quote, error) {
	args := m.Called(category)
	return args.Get(0).(*Quote), args.Error(1)
}

func (m *MockQuoteRepository) Update(ID int, category *Quote) (*Quote, error) {
	args := m.Called(ID, category)
	return args.Get(0).(*Quote), args.Error(1)
}

func (m *MockQuoteRepository) Delete(ID int) error {
	args := m.Called(ID)
	return args.Error(0)
}
