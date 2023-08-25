package usecase

import (
	"log"

	"github.com/nunenuh/iquote-fiber/internal/core/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/core/utils/exception"
	"github.com/nunenuh/iquote-fiber/internal/core/utils/hash"
)

type AuthUsecase struct {
	repo domain.IAuthorRepository
}

func NewAuthUsecase(repo domain.IAuthorRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (ucase *AuthUsecase) Login(username string, password string) (*domain.Login, error) {
	user, err := ucase.repo.GetByUsername(username)
	if err != nil {
		return nil, exception.NewOtherError("Forbidden!")
	}

	log.Printf("username: %s, password: %s", username, user.Password)

	if !hash.CheckHashPassword(password, user.Password) {
		return nil, exception.NewOtherError("Invalid Credentials!")
	}

	return user, nil
}
