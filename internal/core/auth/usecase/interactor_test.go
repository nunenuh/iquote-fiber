package usecase

import (
	"errors"
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/core/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/core/utils/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginSuccessLogin(t *testing.T) {
	var (
		repo    = &domain.MockAuthRepository{Mock: mock.Mock{}}
		svc     = &domain.MockAuthService{Mock: mock.Mock{}}
		usecase = NewAuthUsecase(repo, svc)
	)

	hpass, err := hash.HashPassword("test")
	assert.Nil(t, err)

	repoAuthExpected := domain.Auth{
		Username: "test",
		Password: hpass,
	}

	repo.Mock.On("GetByUsername", "test").Return(repoAuthExpected, nil)
	svc.Mock.On("GenerateToken", repoAuthExpected).Return("token_string", nil)

	token, err := usecase.Login("test", "test")
	assert.Equal(t, token, "token_string")
	assert.Nil(t, err)
}

func TestLoginFailedUserNotFound(t *testing.T) {
	var (
		repo    = &domain.MockAuthRepository{Mock: mock.Mock{}}
		svc     = &domain.MockAuthService{Mock: mock.Mock{}}
		usecase = NewAuthUsecase(repo, svc)
	)
	repo.Mock.On("GetByUsername", "test").Return(nil, errors.New("failed to get user"))
	svc.Mock.On("GenerateToken", mock.Anything).Return("", errors.New("failed to generate token"))

	token, err := usecase.Login("test", "test")
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
}

func TestLoginFailedGenerateTokenFailed(t *testing.T) {
	var (
		repo    = &domain.MockAuthRepository{Mock: mock.Mock{}}
		svc     = &domain.MockAuthService{Mock: mock.Mock{}}
		usecase = NewAuthUsecase(repo, svc)
	)
	repoAuthExpected := domain.Auth{Username: "test", Password: "test"}
	repo.Mock.On("GetByUsername", "test").Return(repoAuthExpected, nil)
	svc.Mock.On("GenerateToken", repoAuthExpected).Return("", errors.New("failed to generate token"))

	token, err := usecase.Login("test", "test")
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
}

func TestVerifyTokenSuccess(t *testing.T) {
	var (
		repo    = &domain.MockAuthRepository{Mock: mock.Mock{}}
		svc     = &domain.MockAuthService{Mock: mock.Mock{}}
		usecase = NewAuthUsecase(repo, svc)
	)

	tokenString := "token_string"
	cClaims := domain.CustomClaims{
		UserID:   1,
		Username: "admin",
		Email:    "test@gmail.com",
	}

	svc.Mock.On("VerifyToken", "token_string").Return(&cClaims, nil)

	cClaimsResult, err := usecase.VerifyToken(tokenString)
	assert.Equal(t, cClaimsResult.Email, cClaims.Email)
	assert.Equal(t, cClaimsResult.UserID, cClaims.UserID)
	assert.Equal(t, cClaimsResult.Username, cClaims.Username)
	assert.Nil(t, err)
}

func TestRefreshTokenSuccess(t *testing.T) {
	var (
		repo    = &domain.MockAuthRepository{Mock: mock.Mock{}}
		svc     = &domain.MockAuthService{Mock: mock.Mock{}}
		usecase = NewAuthUsecase(repo, svc)
	)

	tokenInput := "token_input"
	tokenRefreshed := "refresh_token"

	svc.Mock.On("RefreshToken", tokenInput).Return(tokenRefreshed, nil)
	token, err := usecase.RefreshToken(tokenInput)
	assert.Nil(t, err)
	assert.Equal(t, token, tokenRefreshed)

}
