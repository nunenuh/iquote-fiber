package domain

type IAuthRepository interface {
	GetByUsername(username string) (*Auth, error)
	// GetByEmail(email string) (*Login, error)
}
