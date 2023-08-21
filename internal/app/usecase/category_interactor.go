package usecase

import (
	exception "github.com/nunenuh/iquote-fiber/internal/app/exeption"
	"github.com/nunenuh/iquote-fiber/internal/app/validator"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type CategoryUsecase struct {
	repo      repository.ICategoryRepository
	validator *validator.Validator
}

func NewCategoryUsecase(r repository.ICategoryRepository) *CategoryUsecase {
	validator := validator.NewValidator()
	return &CategoryUsecase{
		repo:      r,
		validator: validator,
	}
}

func (ucase *CategoryUsecase) GetAll(limit int, offset int) ([]*entity.Category, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByID(ID int) (*entity.Category, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByName(name string) ([]*entity.Category, error) {
	u, err := ucase.repo.GetByName(name)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByParentID(ID int) ([]*entity.Category, error) {
	u, err := ucase.repo.GetByParentID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *CategoryUsecase) Create(category *entity.Category) (*entity.Category, error) {
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

func (ucase *CategoryUsecase) Update(ID int, category *entity.Category) (*entity.Category, error) {
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
