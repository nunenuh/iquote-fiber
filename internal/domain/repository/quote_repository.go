package repository

import "github.com/nunenuh/iquote-fiber/internal/domain/entity"

type IQuoteRepository interface {
	GetAll(limit int, offset int) ([]*entity.Quote, error)
	GetByID(ID int) (*entity.Quote, error)
	Create(quote *entity.Quote) (*entity.Quote, error)
	Update(ID int, Quote *entity.Quote) (*entity.Quote, error)
	Delete(ID int) error
}
