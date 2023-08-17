package repository

import "github.com/nunenuh/iquote-fiber/internal/domain/entity"

type IAuthorRepository interface {
	GetAll(limit int, offset int) ([]*entity.Author, error)
	GetByName(name string) (*entity.Author, error)
	GetByID(ID int) (*entity.Author, error)
	Create(author *entity.Author) (*entity.Author, error)
	Update(ID int, author *entity.Author) (*entity.Author, error)
	Delete(ID int) error
}
