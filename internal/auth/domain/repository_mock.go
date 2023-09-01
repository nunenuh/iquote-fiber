package domain

import (
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	Mock mock.Mock
}

func NewMockAuthRepository(mock mock.Mock) *MockAuthRepository {
	return &MockAuthRepository{
		Mock: mock,
	}
}

func (m *MockAuthRepository) GetByUsername(username string) (*Auth, error) {
	args := m.Mock.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		auth := args.Get(0).(Auth)
		return &auth, nil
	}
}
