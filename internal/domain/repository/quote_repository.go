package repository

import "github.com/nunenuh/iquote-fiber/internal/domain/entity"

type IQuoteRepository interface {
	GetAll(limit int, offset int) ([]*entity.Quote, error)
	GetByAuthorName(name string, limit int, offset int) ([]*entity.Quote, error)
	GetByAuthorID(ID int, limit int, offset int) ([]*entity.Quote, error)
	GetByCategoryName(name string, limit int, offset int) ([]*entity.Quote, error)
	GetByCategoryID(ID int, limit int, offset int) ([]*entity.Quote, error)
	// GetByTags(tags string) ([]*entity.Quote, error)
	// Search(keyword string) ([]*entity.Quote, error)
	Like(quoteID int, userID int) (*entity.Quote, error)
	// Unlike(quoteID int, userID int) (*entity.Author, error)
	GetByID(ID int) (*entity.Quote, error)
	Create(quote *entity.Quote) (*entity.Quote, error)
	Update(ID int, quote *entity.Quote) (*entity.Quote, error)
	Delete(ID int) error
}
