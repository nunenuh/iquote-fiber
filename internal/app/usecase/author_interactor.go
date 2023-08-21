package usecase

import (
	exception "github.com/nunenuh/iquote-fiber/internal/app/exeption"
	"github.com/nunenuh/iquote-fiber/internal/app/validator"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type AuthorUsecase struct {
	repo      repository.IAuthorRepository
	validator *validator.Validator
}

func NewAuthorUsecase(r repository.IAuthorRepository) *AuthorUsecase {
	validator := validator.NewValidator()
	return &AuthorUsecase{
		repo:      r,
		validator: validator,
	}
}

func (ucase *AuthorUsecase) GetAll(limit int, offset int) ([]*entity.Author, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *AuthorUsecase) GetByID(ID int) (*entity.Author, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *AuthorUsecase) GetByName(name string) (*entity.Author, error) {
	u, err := ucase.repo.GetByName(name)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *AuthorUsecase) Create(author *entity.Author) (*entity.Author, error) {
	err := ucase.validator.Validate(author)
	if err != nil {
		return nil, exception.NewValidatorError(err.Error())
	}

	u, err := ucase.repo.Create(author)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}
	return u, nil
}

func (ucase *AuthorUsecase) Update(ID int, author *entity.Author) (*entity.Author, error) {
	err := ucase.validator.Validate(author)
	if err != nil {
		return nil, exception.NewValidatorError(err.Error())
	}

	u, err := ucase.repo.Update(ID, author)
	if err != nil {
		return nil, exception.NewRepositoryError(err.Error())
	}

	return u, nil
}

func (ucase *AuthorUsecase) Delete(ID int) error {
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
