package domain

type IUserRepository interface {
	GetAll(limit int, offset int) ([]*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByID(ID int) (*User, error)
	Create(user *User) (*User, error)
	Update(ID int, user *User) (*User, error)
	Delete(ID int) error
}
