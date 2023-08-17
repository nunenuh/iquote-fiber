package repository

type IAuthRepository interface {
	Login(username string, password string) (bool, error)
	RefreshToken(token string) (string, error)
	VerifyToken(token string) (string, error)
}
