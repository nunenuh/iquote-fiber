package repository

import "github.com/nunenuh/iquote-fiber/internal/domain/entity"

type IUserRepository interface {
	GetByID(ID int) (*entity.User, error)
	// GetAll(limit int, offset int) ([]*entity.User, error)
	// Create(user entity.User) (*entity.User, error)
	// Update(user entity.User) (*entity.User, error)
	// Delete(ID int) error
}
