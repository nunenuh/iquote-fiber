package domain

type IAuthorRepository interface {
	GetByUsername(username string) (*Login, error)
	// GetByEmail(email string) (*Login, error)
}
