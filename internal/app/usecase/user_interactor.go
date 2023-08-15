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

func (user *UserUsecase) GetByID(ID int) (*entity.User, error) {
	u, err := user.repo.GetByID(ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
