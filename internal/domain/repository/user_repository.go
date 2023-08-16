package repository

import "github.com/nunenuh/iquote-fiber/internal/domain/entity"

type IUserRepository interface {
	GetAll(limit int, offset int) ([]*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetByID(ID int) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Update(ID int, user *entity.User) (*entity.User, error)
	Delete(ID int) error
}
