package domain

type IQuoteRepository interface {
	GetAll(limit int, offset int) ([]*Quote, error)
	GetByAuthorName(name string, limit int, offset int) ([]*Quote, error)
	GetByAuthorID(ID int, limit int, offset int) ([]*Quote, error)
	GetByCategoryName(name string, limit int, offset int) ([]*Quote, error)
	GetByCategoryID(ID int, limit int, offset int) ([]*Quote, error)
	// GetByTags(tags string) ([]*Quote, error)
	// Search(keyword string) ([]*Quote, error)
	Like(quoteID int, userID int) (*Quote, error)
	Unlike(quoteID int, userID int) (*Quote, error)
	GetByID(ID int) (*Quote, error)
	Create(quote *Quote) (*Quote, error)
	Update(ID int, quote *Quote) (*Quote, error)
	Delete(ID int) error
}
