package usecase

import (
	"errors"

	"github.com/nunenuh/iquote-fiber/internal/app/utils"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type AuthUsecase struct {
	repo repository.IUserRepository
}

func NewAuthUsecase(repo repository.IUserRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (ucase *AuthUsecase) Login(username string, password string) (*entity.User, error) {
	user, err := ucase.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("Forbidden!")
	}

	if !utils.CheckHashPassword(password, user.Password) {
		return nil, errors.New("Invalid Credentials!")
	}

	return user, nil
}
