package usecase

import (
	"errors"
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/shared/exception"
	"github.com/stretchr/testify/assert"
)

func TestCategoryUsecaseGetAll(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	expectedCategorys := []*domain.Category{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 1"},
	}

	mockRepo.On("GetAll", 10, 0).Return(expectedCategorys, nil)

	categorys, err := uc.GetAll(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategorys, categorys)
}

func TestCategoryUsecaseGetByID(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	expectedCategory := &domain.Category{ID: 1, Name: "Category1"}

	mockRepo.On("GetByID", 1).Return(expectedCategory, nil)

	category, err := uc.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestCategoryUsecaseGetByName(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)
	expectedCategory := []*domain.Category{
		{Name: "myname", Description: "myname"},
	}

	mockRepo.On("GetByName", "myname").Return(expectedCategory, nil)

	category, err := uc.GetByName("myname")
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestCategoryUsecaseGetByParentID(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)
	expectedCategory := []*domain.Category{
		{ID: 3, Name: "myname 1", Description: "myname 1", ParentID: 1},
		{ID: 4, Name: "myname 2", Description: "myname 2", ParentID: 1},
	}
	mockRepo.On("GetByParentID", 1).Return(expectedCategory, nil)

	category, err := uc.GetByParentID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, category)
}

func TestCategoryUsecaseCreate(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	newCategory := &domain.Category{
		Name:        "Category1",
		Description: "description",
	}

	mockRepo.On("Create", newCategory).Return(newCategory, nil)

	category, err := uc.Create(newCategory)
	assert.NoError(t, err)
	assert.Equal(t, newCategory, category)
}

func TestCategoryUsecaseCreateErrorValidation(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	newCategory := &domain.Category{} // This category should fail validation

	category, err := uc.Create(newCategory)
	assert.Error(t, err)
	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.ValidatorError, appErr.Type)

	assert.Nil(t, category)
}

func TestCategoryUsecaseCreateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	newCategory := &domain.Category{
		Name:        "Category1",
		Description: "description",
	}

	mockRepo.On("Create", newCategory).Return((*domain.Category)(nil), errors.New("create error"))

	category, err := uc.Create(newCategory)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, category)
}

func TestCategoryUsecaseUpdate(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	updatedCategory := &domain.Category{
		ID:          1,
		Name:        "Updated-Category",
		Description: "updated description",
	}

	mockRepo.On("Update", 1, updatedCategory).Return(updatedCategory, nil)

	category, err := uc.Update(1, updatedCategory)
	assert.NoError(t, err)
	assert.Equal(t, updatedCategory, category)
}

func TestCategoryUsecaseUpdateErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	updatedCategory := &domain.Category{ID: 1, Name: "Updated Category", Description: "updated description"}

	mockRepo.On("Update", 1, updatedCategory).Return((*domain.Category)(nil), errors.New("update error"))

	category, err := uc.Update(1, updatedCategory)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)

	assert.Nil(t, category)
}

func TestCategoryUsecaseUpdateErrorValidation(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	updatedCategory := &domain.Category{ID: 1, Name: "Updated Category"}

	mockRepo.On("Update", 1, updatedCategory).Return((*domain.Category)(nil), errors.New("update error"))

	category, err := uc.Update(1, updatedCategory)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.ValidatorError, appErr.Type)

	assert.Nil(t, category)
}

func TestCategoryUsecaseDelete(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return(&domain.Category{ID: 1}, nil)
	mockRepo.On("Delete", 1).Return(nil)

	err := uc.Delete(1)
	assert.NoError(t, err)
}

func TestCategoryUsecaseDeleteErrorRepository(t *testing.T) {
	mockRepo := new(domain.MockCategoryRepository)
	uc := NewCategoryUsecase(mockRepo)

	mockRepo.On("GetByID", 1).Return((*domain.Category)(nil), errors.New("not found"))
	mockRepo.On("Delete", 1).Return(errors.New("not found"))

	err := uc.Delete(1)
	assert.Error(t, err)

	// Check if the error is of the desired type
	appErr, ok := err.(exception.AppError)
	assert.True(t, ok)
	assert.Equal(t, exception.RepositoryError, appErr.Type)
}

// Continue with other tests...
