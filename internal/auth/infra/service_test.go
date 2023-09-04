package infra

import (
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/auth/domain"
	"github.com/nunenuh/iquote-fiber/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_GenerateToken(t *testing.T) {
	cfg := config.Configuration{} // Define your test configuration here
	authService := NewAuthService(cfg)

	auth := domain.Auth{} // Define your test auth object here

	token, err := authService.GenerateToken(auth)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestAuthService_VerifyToken(t *testing.T) {
	cfg := config.Configuration{} // Define your test configuration here
	authService := NewAuthService(cfg)

	tokenString := "" // Define your test token string here

	claims, err := authService.VerifyToken(tokenString)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestAuthService_RefreshToken(t *testing.T) {
	cfg := config.Configuration{} // Define your test configuration here
	authService := NewAuthService(cfg)

	tokenString := "" // Define your test token string here

	newToken, err := authService.RefreshToken(tokenString)

	assert.Error(t, err)
	assert.Empty(t, newToken)
}

func TestAuthService_ParseToken(t *testing.T) {
	cfg := config.Configuration{} // Define your test configuration here
	authService := NewAuthService(cfg)

	auth := domain.Auth{} // Define your test auth object here

	tokenstr, err := authService.GenerateToken(auth)
	assert.NoError(t, err)

	token, err := authService.ParseToken(tokenstr)

	assert.NoError(t, err)
	assert.NotNil(t, token)
}
