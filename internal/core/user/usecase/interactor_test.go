package usecase

import (
	"errors"
	"testing"

	// "github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/core/user/domain"
	// "github.com/nunenuh/iquote-fiber/internal/core/user/usecase"
	"github.com/stretchr/testify/assert"
)

type UserRepositoryMock struct {
	GetByIDFunc       func(ID int) (*domain.User, error)
	GetByUsernameFunc func(username string) (*domain.User, error)
	GetByEmailFunc    func(email string) (*domain.User, error)
	GetAllFunc        func(limit int, offset int) ([]*domain.User, error)
	CreateFunc        func(user *domain.User) (*domain.User, error)
	UpdateFunc        func(ID int, user *domain.User) (*domain.User, error)
	DeleteFunc        func(ID int) error
}

func (m *UserRepositoryMock) GetByID(ID int) (*domain.User, error) {
	return m.GetByIDFunc(ID)
}

func (m *UserRepositoryMock) GetByUsername(username string) (*domain.User, error) {
	return m.GetByUsernameFunc(username)
}

func (m *UserRepositoryMock) GetByEmail(email string) (*domain.User, error) {
	return m.GetByEmailFunc(email)
}

func (m *UserRepositoryMock) GetAll(limit int, offset int) ([]*domain.User, error) {
	return m.GetAllFunc(limit, offset)
}

func (m *UserRepositoryMock) Create(user *domain.User) (*domain.User, error) {
	return m.CreateFunc(user)
}

func (m *UserRepositoryMock) Update(ID int, user *domain.User) (*domain.User, error) {
	return m.UpdateFunc(ID, user)
}

func (m *UserRepositoryMock) Delete(ID int) error {
	return m.DeleteFunc(ID)
}

func TestUserUsecase_GetByID(t *testing.T) {
	t.Run("User found", func(t *testing.T) {
		expectedUser := &domain.User{ID: 1, FullName: "John Doe"}
		repo := &UserRepositoryMock{
			GetByIDFunc: func(ID int) (*domain.User, error) {
				return expectedUser, nil
			},
		}
		usecase := NewUserUsecase(repo)

		user, err := usecase.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("User not found", func(t *testing.T) {
		repo := &UserRepositoryMock{
			GetByIDFunc: func(ID int) (*domain.User, error) {
				return nil, errors.New("user not found")
			},
		}
		usecase := NewUserUsecase(repo)

		user, err := usecase.GetByID(1)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
	})
}
