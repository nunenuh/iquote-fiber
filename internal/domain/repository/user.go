package repository

import "github.com/nunenuh/iquote-fiber/internal/domain"

type IUserRepository interface {
	GetByID(ID int) domain.User
	GetAll(limit int, offset int) []domain.User
	Create(user domain.User) domain.User
	Update(user domain.User) domain.User
	Delete(ID int)
}
