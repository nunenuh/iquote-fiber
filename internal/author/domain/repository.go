package domain

type IAuthorRepository interface {
	GetAll(limit int, offset int) ([]*Author, error)
	GetByName(name string) (*Author, error)
	GetByID(ID int) (*Author, error)
	Create(author *Author) (*Author, error)
	Update(ID int, author *Author) (*Author, error)
	Delete(ID int) error
}
