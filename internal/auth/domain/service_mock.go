package domain

import (
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	Mock mock.Mock
}

func NewMockAuthService(mock mock.Mock) *MockAuthService {
	return &MockAuthService{
		Mock: mock,
	}
}

func (m *MockAuthService) GenerateToken(auth Auth) (string, error) {
	args := m.Mock.Called(auth)
	if args.Get(0) == nil {
		return "", args.Error(1)
	} else {
		token := args.Get(0).(string)
		return token, nil
	}
}

func (m *MockAuthService) RefreshToken(tokenString string) (string, error) {
	args := m.Mock.Called(tokenString)
	if args.Get(0) == nil {
		return "", args.Error(1)
	} else {
		token := args.Get(0).(string)
		return token, nil
	}
}

func (m *MockAuthService) VerifyToken(tokenString string) (*CustomClaims, error) {
	args := m.Mock.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		claims := args.Get(0).(*CustomClaims)
		return claims, nil
	}
}
