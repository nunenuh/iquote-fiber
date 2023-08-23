package usecase

import (
	"log"

	exception "github.com/nunenuh/iquote-fiber/internal/app/exeption"
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
		return nil, exception.NewOtherError("Forbidden!")
	}

	log.Printf("username: %s, password: %s", username, user.Password)

	// passwd, err := utils.HashPassword(password)
	// if err != nil {
	// 	return nil, exception.NewOtherError(err.Error())
	// }

	if !utils.CheckHashPassword(password, user.Password) {
		return nil, exception.NewOtherError("Invalid Credentials!")
	}

	return user, nil
}
