package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/domain"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

func User(r repository.IUserRepository) domain.User {
	return r.Get()
}
