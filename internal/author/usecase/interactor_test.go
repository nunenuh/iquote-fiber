package usecase

import (
	"errors"
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/author/domain"
	"github.com/nunenuh/iquote-fiber/internal/shared/exception"
	"github.com/stretchr/testify/assert"
)

// func TestGetAll(t *testing.T) {
// 	mockRepo := new(domain.MockAuthorRepository)
// 	uc := NewAuthorUsecase(mockRepo)

// 	expectedAuthors := []*domain.Author{
// 		{ID: 1, Name: "Author1"},
// 		{ID: 2, Name: "Author2"},
// 	}

// 	mockRepo.On("GetAll", 10, 0).Return(expectedAuthors, nil)

// 	authors, err := uc.GetAll(10, 0)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedAuthors, authors)
// }

func TestGetByID(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	expectedAuthor := &domain.Author{ID: 1, Name: "Author1"}

	mockRepo.On("GetByID", 1).Return(expectedAuthor, nil)

	author, err := uc.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthor, author)
}

func TestGetByName(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	expectedAuthor := &domain.Author{ID: 1, Name: "Author1"}

	mockRepo.On("GetByName", "myname").Return(expectedAuthor, nil)

	author, err := uc.GetByName("myname")

	assert.NoError(t, err)
	assert.Equal(t, expectedAuthor, author)
}

func TestCreate(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	newAuthor := &domain.Author{
		Name:        "Author1",
		Description: "description",
	}

	mockRepo.On("Create", newAuthor).Return(newAuthor, nil)

	author, err := uc.Create(newAuthor)
	assert.NoError(t, err)
	assert.Equal(t, newAuthor, author)
}

func TestCreateErrorValidation(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	newAuthor := &domain.Author{} // This author should fail validation

	author, err := uc.Create(newAuthor)
	assert.Error(t, err)
	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.ValidatorError, appErr.Type)

	assert.Nil(t, author)
}

func TestCreateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	newAuthor := &domain.Author{
		Name:        "Author1",
		Description: "description",
	}

	mockRepo.On("Create", newAuthor).Return((*domain.Author)(nil), errors.New("create error"))

	author, err := uc.Create(newAuthor)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, author)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	updatedAuthor := &domain.Author{ID: 1, Name: "Updated Author", Description: "updated description"}

	mockRepo.On("Update", 1, updatedAuthor).Return(updatedAuthor, nil)

	author, err := uc.Update(1, updatedAuthor)
	assert.NoError(t, err)
	assert.Equal(t, updatedAuthor, author)
}

func TestUpdateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	updatedAuthor := &domain.Author{ID: 1, Name: "Updated Author", Description: "updated description"}

	mockRepo.On("Update", 1, updatedAuthor).Return((*domain.Author)(nil), errors.New("update error"))

	author, err := uc.Update(1, updatedAuthor)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, author)
}

func TestUpdateErrorValidation(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	updatedAuthor := &domain.Author{ID: 1, Name: "Updated Author"}

	mockRepo.On("Update", 1, updatedAuthor).Return((*domain.Author)(nil), errors.New("update error"))

	author, err := uc.Update(1, updatedAuthor)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.ValidatorError, appErr.Type)

	assert.Nil(t, author)
}

func TestDelete(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return(&domain.Author{ID: 1}, nil)
	mockRepo.On("Delete", 1).Return(nil)

	err := uc.Delete(1)
	assert.NoError(t, err)
}

func TestDeleteErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockAuthorRepository)
	uc := NewAuthorUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return((*domain.Author)(nil), errors.New("not found"))
	mockRepo.On("Delete", 1).Return(errors.New("not found"))

	err := uc.Delete(1)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)
}

// Continue with other tests...
