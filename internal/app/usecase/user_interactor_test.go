package usecase_test

import (
	"errors"
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type UserRepositoryMock struct {
	GetByIDFunc func(ID int) (*entity.User, error)
}

func (m *UserRepositoryMock) GetByID(ID int) (*entity.User, error) {
	return m.GetByIDFunc(ID)
}

func TestUserUsecase_GetByID(t *testing.T) {
	t.Run("User found", func(t *testing.T) {
		expectedUser := &entity.User{ID: 1, FullName: "John Doe"}
		repo := &UserRepositoryMock{
			GetByIDFunc: func(ID int) (*entity.User, error) {
				return expectedUser, nil
			},
		}
		usecase := usecase.NewUserUsecase(repo)

		user, err := usecase.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("User not found", func(t *testing.T) {
		repo := &UserRepositoryMock{
			GetByIDFunc: func(ID int) (*entity.User, error) {
				return nil, errors.New("user not found")
			},
		}
		usecase := usecase.NewUserUsecase(repo)

		user, err := usecase.GetByID(1)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
	})
}
