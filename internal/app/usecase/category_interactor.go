package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type CategoryUsecase struct {
	repo repository.ICategoryRepository
}

func NewCategoryUsecase(r repository.ICategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		repo: r,
	}
}

func (ucase *CategoryUsecase) GetAll(limit int, offset int) ([]*entity.Category, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByID(ID int) (*entity.Category, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByName(name string) ([]*entity.Category, error) {
	u, err := ucase.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *CategoryUsecase) GetByParentID(ID int) ([]*entity.Category, error) {
	u, err := ucase.repo.GetByParentID(ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *CategoryUsecase) Create(author *entity.Category) (*entity.Category, error) {
	u, err := ucase.repo.Create(author)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ucase *CategoryUsecase) Update(ID int, author *entity.Category) (*entity.Category, error) {
	u, err := ucase.repo.Update(ID, author)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *CategoryUsecase) Delete(ID int) error {
	err := ucase.repo.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
