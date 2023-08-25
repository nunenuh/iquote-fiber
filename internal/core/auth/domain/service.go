package domain

type IAuthService interface {
	GenerateToken(auth Auth) (string, error)
	VerifyToken(tokenString string) (*CustomClaims, error)
	RefreshToken(tokenString string) (string, error)
}
