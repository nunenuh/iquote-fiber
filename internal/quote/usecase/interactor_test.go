package usecase

import (
	"errors"
	"testing"

	quote_domain "github.com/nunenuh/iquote-fiber/internal/author/domain"
	category_domain "github.com/nunenuh/iquote-fiber/internal/category/domain"

	"github.com/nunenuh/iquote-fiber/internal/quote/domain"
	"github.com/nunenuh/iquote-fiber/internal/shared/exception"
	"github.com/stretchr/testify/assert"
)

func TestQuoteUsecaseGetAll(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	expectedQuotes := []*domain.Quote{
		{ID: 1, QText: "Quote1"},
		{ID: 2, QText: "Quote2"},
	}

	mockRepo.On("GetAll", 10, 0).Return(expectedQuotes, nil)

	quotes, err := uc.GetAll(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuotes, quotes)
}

func TestQuoteUsecaseGetByID(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	expectedQuote := &domain.Quote{
		ID:     1,
		QText:  "Quote1",
		Tags:   "",
		Author: quote_domain.Author{ID: 1, Name: "Author1"},
	}

	mockRepo.On("GetByID", 1).Return(expectedQuote, nil)

	quote, err := uc.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuote, quote)
}

func TestQuoteUsecaseGetByAuthorName(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	expectedQuote := []*domain.Quote{
		{ID: 1, QText: "myname", Tags: "myname"},
	}

	limit := 10
	offset := 0
	authorName := "author1"

	mockRepo.On("GetByAuthorName", authorName, limit, offset).Return(expectedQuote, nil)

	quote, err := uc.GetByAuthorName(authorName, limit, offset)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuote, quote)
}

func TestQuoteUsecaseGetByAuthorID(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	expectedQuote := []*domain.Quote{
		{
			QText: "myname", Tags: "myname",
			Author: quote_domain.Author{ID: 1, Name: "Author1"},
		},
	}

	limit := 10
	offset := 0
	authorID := 2

	mockRepo.On("GetByAuthorID", authorID, limit, offset).Return(expectedQuote, nil)

	quote, err := uc.GetByAuthorID(authorID, limit, offset)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuote, quote)
}

func TestQuoteUsecaseGetByCategoryName(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	expectedQuote := []*domain.Quote{
		{
			ID: 1, QText: "myname", Tags: "myname",
			Category: []category_domain.Category{
				{ID: 1, Name: "cat"},
			},
		},
	}

	limit := 10
	offset := 0
	categoryName := "cat"

	mockRepo.On("GetByCategoryName", categoryName, limit, offset).Return(expectedQuote, nil)

	quote, err := uc.GetByCategoryName(categoryName, limit, offset)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuote, quote)
}

func TestQuoteUsecaseGetByCategoryID(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	expectedQuote := []*domain.Quote{
		{
			QText: "myname", Tags: "myname",
			Author: quote_domain.Author{ID: 1, Name: "Author1"},
			Category: []category_domain.Category{
				{ID: 1, Name: "cat"},
			},
		},
	}

	limit := 10
	offset := 0
	categoryID := 1

	mockRepo.On("GetByCategoryID", categoryID, limit, offset).Return(expectedQuote, nil)

	quote, err := uc.GetByCategoryID(categoryID, limit, offset)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuote, quote)
}

func TestQuoteUsecaseCreate(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	newQuote := &domain.Quote{
		QText:  "Quote1",
		Tags:   "description",
		Author: quote_domain.Author{ID: 1, Name: "Author1"},
	}

	mockRepo.On("Create", newQuote).Return(newQuote, nil)

	quote, err := uc.Create(newQuote)
	assert.NoError(t, err)
	assert.Equal(t, newQuote, quote)
}

func TestQuoteUsecaseCreateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	newQuote := &domain.Quote{
		ID:     1,
		QText:  "Quote1",
		Tags:   "description",
		Author: quote_domain.Author{ID: 1, Name: "Author1"},
	}

	mockRepo.On("Create", newQuote).Return((*domain.Quote)(nil), errors.New("create error"))

	quote, err := uc.Create(newQuote)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, quote)
}

func TestQuoteUsecaseUpdate(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	updatedQuote := &domain.Quote{
		ID:     1,
		QText:  "Quote1",
		Tags:   "description",
		Author: quote_domain.Author{ID: 1, Name: "Author1"},
	}

	mockRepo.On("Update", 1, updatedQuote).Return(updatedQuote, nil)

	quote, err := uc.Update(1, updatedQuote)
	assert.NoError(t, err)
	assert.Equal(t, updatedQuote, quote)
}

func TestQuoteUsecaseUpdateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	updatedQuote := &domain.Quote{
		ID:     1,
		QText:  "Quote1",
		Tags:   "description",
		Author: quote_domain.Author{ID: 1, Name: "Author1"},
	}

	mockRepo.On("Update", 1, updatedQuote).Return((*domain.Quote)(nil), errors.New("update error"))

	quote, err := uc.Update(1, updatedQuote)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, quote)
}

func TestQuoteUsecaseDelete(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return(&domain.Quote{ID: 1}, nil)
	mockRepo.On("Delete", 1).Return(nil)

	err := uc.Delete(1)
	assert.NoError(t, err)
}

func TestQuoteUsecaseDeleteErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockQuoteRepository)
	uc := NewQuoteUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return((*domain.Quote)(nil), errors.New("not found"))
	mockRepo.On("Delete", 1).Return(errors.New("not found"))

	err := uc.Delete(1)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)
}

// Continue with other tests...
