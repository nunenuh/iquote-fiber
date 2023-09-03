package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/shared/exception"
	"github.com/nunenuh/iquote-fiber/internal/shared/validator"
)

type CategoryUsecase struct {
	repo      domain.ICategoryRepository
	validator *validator.Validator
}

func NewCategoryUsecase(r domain.ICategoryRepository) *CategoryUsecase {
	validator := validator.NewValidator()
	return &CategoryUsecase{
		repo:      r,
		validator: validator,
	}
}

func (ucase *CategoryUsecase) GetAll(limit int, offset int) ([]*domain.Category, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByID(ID int) (*domain.Category, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByName(name string) ([]*domain.Category, error) {
	u, err := ucase.repo.GetByName(name)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByParentID(ID int) ([]*domain.Category, error) {
	u, err := ucase.repo.GetByParentID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) Create(category *domain.Category) (*domain.Category, error) {
	err := ucase.validator.Validate(category)
	if err != nil {
		return nil, exception.NewValidatorError(err.Error())
	}

	u, err := ucase.repo.Create(category)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}
	return u, nil
}

func (ucase *CategoryUsecase) Update(ID int, category *domain.Category) (*domain.Category, error) {
	err := ucase.validator.Validate(category)
	if err != nil {
		return nil, exception.NewValidatorError(err.Error())
	}

	u, err := ucase.repo.Update(ID, category)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) Delete(ID int) error {
	_, errGet := ucase.repo.GetByID(ID)
	if errGet != nil {
		return exception.NewRepositoryError(errGet.Error())
	}

	err := ucase.repo.Delete(ID)
	if err != nil {
		return exception.NewRepositoryError(err.Error())
	}

	return nil
}
