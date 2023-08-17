package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type AuthorUsecase struct {
	repo repository.IAuthorRepository
}

func NewAuthorUsecase(r repository.IAuthorRepository) *AuthorUsecase {
	return &AuthorUsecase{
		repo: r,
	}
}

func (ucase *AuthorUsecase) GetAll(limit int, offset int) ([]*entity.Author, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *AuthorUsecase) GetByID(ID int) (*entity.Author, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *AuthorUsecase) Create(author *entity.Author) (*entity.Author, error) {
	u, err := ucase.repo.Create(author)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ucase *AuthorUsecase) Update(ID int, author *entity.Author) (*entity.Author, error) {
	u, err := ucase.repo.Update(ID, author)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *AuthorUsecase) Delete(ID int) error {
	err := ucase.repo.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
