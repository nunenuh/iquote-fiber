package repository

import "github.com/nunenuh/iquote-fiber/internal/domain/entity"

type ICategoryRepository interface {
	GetAll(limit int, offset int) ([]*entity.Category, error)
	GetByName(name string) (*entity.Category, error)
	GetByParentID(ID int) ([]*entity.Category, error)
	GetByID(ID int) (*entity.Category, error)
	Create(category *entity.Category) (*entity.Category, error)
	Update(ID int, category *entity.Category) (*entity.Category, error)
	Delete(ID int) error
}
