package usecase

import (
	"github.com/nunenuh/iquote-fiber/internal/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/utils/exception"
	"github.com/nunenuh/iquote-fiber/internal/utils/hash"
)

type AuthUsecase struct {
	repo domain.IAuthRepository
	svc  domain.IAuthService
}

func NewAuthUsecase(repo domain.IAuthRepository, svc domain.IAuthService) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
		svc:  svc,
	}
}

func (ucase *AuthUsecase) Login(username string, password string) (string, error) {
	auth, err := ucase.repo.GetByUsername(username)
	if err != nil {
		return "", exception.NewOtherError("Forbidden!")
	}

	if !hash.CheckHashPassword(password, auth.Password) {
		return "", exception.NewOtherError("Invalid Credentials!")
	}

	tokenString, err := ucase.svc.GenerateToken(*auth)
	if err != nil {
		return "", exception.NewServiceError("Failed to generate token!")
	}

	return tokenString, nil
}

func (ucase *AuthUsecase) RefreshToken(token string) (string, error) {
	tokenString, err := ucase.svc.RefreshToken(token)
	if err != nil {
		return "", exception.NewServiceError("Failed to refresh token!")
	}

	return tokenString, nil
}

func (ucase *AuthUsecase) VerifyToken(tokenString string) (*domain.CustomClaims, error) {
	cClaims, err := ucase.svc.VerifyToken(tokenString)
	if err != nil {
		return nil, exception.NewServiceError("Verify token failed!")
	}

	return cClaims, nil
}
