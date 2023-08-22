package repository

import (
	"fmt"
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type mockUserRepository struct {
	GetAllFunc        func(limit int, offset int) ([]*entity.User, error)
	GetByIDFunc       func(ID int) (*entity.User, error)
	GetByUsernameFunc func(username string) (*entity.User, error)
	GetByEmailFunc    func(email string) (*entity.User, error)
	CreateFunc        func(user *entity.User) (*entity.User, error)
	UpdateFunc        func(ID int, user *entity.User) (*entity.User, error)
	DeleteFunc        func(ID int) error
}

func (m *mockUserRepository) GetAll(limit int, offset int) ([]*entity.User, error) {
	return m.GetAllFunc(limit, offset)
}

func (m *mockUserRepository) GetByID(ID int) (*entity.User, error) {
	return m.GetByIDFunc(ID)
}

func (m *mockUserRepository) GetByUsername(username string) (*entity.User, error) {
	return m.GetByUsernameFunc(username)
}
func (m *mockUserRepository) GetByEmail(email string) (*entity.User, error) {
	return m.GetByEmailFunc(email)
}

func (m *mockUserRepository) Create(user *entity.User) (*entity.User, error) {
	return m.CreateFunc(user)
}

func (m *mockUserRepository) Update(ID int, user *entity.User) (*entity.User, error) {
	return m.UpdateFunc(ID, user)
}

func (m *mockUserRepository) Delete(ID int) error {
	return m.DeleteFunc(ID)
}

// Implement mock functions for the other methods if needed

func TestIUserRepository_GetByID(t *testing.T) {
	// Create a new instance of the mockUserRepository
	mock := &mockUserRepository{}

	// Set the mock behavior for the GetByID method
	expectedUser := &entity.User{ID: 1, FullName: "John Doe"}
	mock.GetByIDFunc = func(ID int) (*entity.User, error) {
		x := expectedUser.ID
		if x != ID {
			return nil, fmt.Errorf("user not found")
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
