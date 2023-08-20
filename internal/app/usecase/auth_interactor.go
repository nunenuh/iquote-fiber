package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type AuthUsecase struct {
	repo repository.IAuthRepository
}

func NewAuthUsecase(r repository.IAuthRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: r,
	}
}

func (ucase *AuthUsecase) Login(username string, password string) (bool, error) {
	result, err := ucase.repo.Login(username, password)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (ucase *AuthUsecase) RefreshToken(token string) (string, error) {
	result, err := ucase.repo.RefreshToken(token)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (ucase *AuthUsecase) VerifyToken(token string) (string, error) {
	result, err := ucase.repo.VerifyToken(token)
	if err != nil {
		return "", err
	}

	return result, nil
}
