package usecase

import (
	"errors"
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/core/user/domain"
	"github.com/nunenuh/iquote-fiber/internal/core/utils/exception"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecaseGetAll(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	expectedUsers := []*domain.User{
		{ID: 1, Username: "user1", FullName: "NameUser1", Email: "johndoe@example.com", IsActive: true},
		{ID: 2, Username: "user2", FullName: "NameUser2", Email: "johndoe@example.com", IsActive: true},
	}

	mockRepo.On("GetAll", 10, 0).Return(expectedUsers, nil)

	users, err := uc.GetAll(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestUserUsecaseGetByID(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	expectedUser := &domain.User{ID: 1, FullName: "User1"}

	mockRepo.On("GetByID", 1).Return(expectedUser, nil)

	user, err := uc.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserUsecaseGetByEmail(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	emailTest := "test@gmail.com"
	expectedUser := &domain.User{
		FullName: "myname",
		Username: "myname",
		Email:    emailTest,
	}

	mockRepo.On("GetByEmail", emailTest).Return(expectedUser, nil)

	user, err := uc.GetByEmail(emailTest)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserUsecaseGetByUsername(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	usernameTest := "username1"
	expectedUser := &domain.User{
		FullName: "myname",
		Username: usernameTest,
		Email:    "test@gmail.com",
	}

	mockRepo.On("GetByUsername", usernameTest).Return(expectedUser, nil)

	user, err := uc.GetByUsername(usernameTest)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserUsecaseCreate(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	newUser := &domain.User{
		FullName: "User1",
		Username: "description",
		Email:    "test2@gmail.com",
	}

	mockRepo.On("Create", newUser).Return(newUser, nil)

	user, err := uc.Create(newUser)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user)
}

func TestUserUsecaseCreateErrorValidation(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	newUser := &domain.User{} // This user should fail validation

	user, err := uc.Create(newUser)
	assert.Error(t, err)
	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.ValidatorError, appErr.Type)

	assert.Nil(t, user)
}

func TestUserUsecaseCreateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	newUser := &domain.User{
		FullName: "User1",
		Username: "description",
		Email:    "test2@gmail.com",
	}

	mockRepo.On("Create", newUser).Return((*domain.User)(nil), errors.New("create error"))

	user, err := uc.Create(newUser)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, user)
}

func TestUserUsecaseUpdate(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	updatedUser := &domain.User{
		ID:       1,
		FullName: "User1",
		Username: "user1",
		Email:    "test2@gmail.com",
	}

	mockRepo.On("Update", 1, updatedUser).Return(updatedUser, nil)

	user, err := uc.Update(1, updatedUser)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, user)
}

func TestUserUsecaseUpdateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	updatedUser := &domain.User{
		ID:       1,
		FullName: "User1",
		Username: "user1",
		Email:    "test2@gmail.com",
	}

	mockRepo.On("Update", 1, updatedUser).Return((*domain.User)(nil), errors.New("update error"))

	user, err := uc.Update(1, updatedUser)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, user)
}

func TestUserUsecaseUpdateErrorValidation(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	updatedUser := &domain.User{}
	mockRepo.On("Update", 1, updatedUser).Return((*domain.User)(nil), errors.New("update error"))

	user, err := uc.Update(1, updatedUser)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.ValidatorError, appErr.Type)

	assert.Nil(t, user)
}

func TestUserUsecaseDelete(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return(&domain.User{ID: 1}, nil)
	mockRepo.On("Delete", 1).Return(nil)

	err := uc.Delete(1)
	assert.NoError(t, err)
}

func TestUserUsecaseDeleteErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockUserRepository)
	uc := NewUserUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return((*domain.User)(nil), errors.New("not found"))
	mockRepo.On("Delete", 1).Return(errors.New("not found"))

	err := uc.Delete(1)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)
}

// Continue with other tests...
