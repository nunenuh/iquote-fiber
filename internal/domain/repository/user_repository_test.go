package repository

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type mockUserRepository struct {
	GetByIDFunc func(ID int) (*entity.User, error)
	// Add mock functions for the other methods if needed
}

func (m *mockUserRepository) GetByID(ID int) (*entity.User, error) {
	return m.GetByIDFunc(ID)
}

func (m *mockUserRepository) GetAll(limit int, offset int) ([]*entity.User, error) {
	return m.GetAll(limit, offset)
}

func (m *mockUserRepository) Create(user *entity.User) (*entity.User, error) {
	return m.Create(user)
}

func (m *mockUserRepository) Update(ID int, user *entity.User) (*entity.User, error) {
	return m.Update(ID, user)
}

func (m *mockUserRepository) Delete(ID int) error {
	return m.Delete(ID)
}

// Implement mock functions for the other methods if needed

func TestIUserRepository_GetByID(t *testing.T) {
	// Create a new instance of the mockUserRepository
	mock := &mockUserRepository{}

	// Set the mock behavior for the GetByID method
	expectedUser := &entity.User{ID: "1", FullName: "John Doe"}
	mock.GetByIDFunc = func(ID int) (*entity.User, error) {
		x, err := strconv.Atoi(expectedUser.ID)
		if ID == x && err == nil {
			return expectedUser, nil
		}
		return nil, fmt.Errorf("user not found")
	}

	// Create an instance of the IUserRepository interface using the mock
	repository := IUserRepository(mock)

	// Call the GetByID method and assert the result
	user, err := repository.GetByID(1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if user != expectedUser {
		t.Errorf("expected user: %v, got: %v", expectedUser, user)
	}
}

// Write additional unit tests for the other methods if needed
