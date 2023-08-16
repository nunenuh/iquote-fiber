package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type UserUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(r repository.IUserRepository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (ucase *UserUsecase) GetAll(limit int, offset int) ([]*entity.User, error) {
	u, err := ucase.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *UserUsecase) GetByID(ID int) (*entity.User, error) {
	u, err := ucase.repo.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *UserUsecase) GetByUsername(username string) (*entity.User, error) {
	u, err := ucase.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *UserUsecase) GetByEmail(email string) (*entity.User, error) {
	u, err := ucase.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *UserUsecase) Create(user *entity.User) (*entity.User, error) {
	u, err := ucase.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ucase *UserUsecase) Update(ID int, user *entity.User) (*entity.User, error) {
	u, err := ucase.repo.Update(ID, user)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ucase *UserUsecase) Delete(ID int) error {
	err := ucase.repo.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
