package domain

import (
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (m *MockUserRepository) GetAll(limit int, offset int) ([]*User, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]*User), args.Error(1)
}

func (m *MockUserRepository) GetByID(ID int) (*User, error) {
	args := m.Called(ID)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(username string) (*User, error) {
	args := m.Called(username)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Create(category *User) (*User, error) {
	args := m.Called(category)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Update(ID int, category *User) (*User, error) {
	args := m.Called(ID, category)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Delete(ID int) error {
	args := m.Called(ID)
	return args.Error(0)
}
