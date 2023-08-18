package repository

import "github.com/nunenuh/iquote-fiber/internal/domain/entity"

type IQuoteRepository interface {
	GetAll(limit int, offset int) ([]*entity.Quote, error)
	// GetByAuthor(name string) ([]*entity.Quote, error)
	// GetByCategory(category string) ([]*entity.Quote, error)
	// GetByTags(tags string) ([]*entity.Quote, error)
	// Search(keyword string) ([]*entity.Quote, error)
	// Like(quoteID int, userID int) (*entity.Quote, error)
	// Unlike(quoteID int, userID int) (*entity.Author, error)
	GetByID(ID int) (*entity.Quote, error)
	Create(category *entity.Quote) (*entity.Quote, error)
	Update(ID int, category *entity.Quote) (*entity.Quote, error)
	Delete(ID int) error
}
