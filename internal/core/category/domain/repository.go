package domain

type ICategoryRepository interface {
	GetAll(limit int, offset int) ([]*Category, error)
	GetByName(name string) ([]*Category, error)
	GetByParentID(ID int) ([]*Category, error)
	GetByID(ID int) (*Category, error)
	Create(category *Category) (*Category, error)
	Update(ID int, category *Category) (*Category, error)
	Delete(ID int) error
}
