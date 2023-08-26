package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	auth := Auth{
		ID:          1,
		Email:       "test@example.com",
		Password:    "password",
		Username:    "testuser",
		FullName:    "Test User",
		Phone:       "1234567890",
		IsSuperuser: false,
	}

	assert.Equal(t, 1, auth.ID, "Expected ID to be 1")
	assert.Equal(t, "test@example.com", auth.Email, "Expected Email to be test@example.com")
	assert.Equal(t, "password", auth.Password, "Expected Password to be password")
	assert.Equal(t, "testuser", auth.Username, "Expected Username to be testuser")
	assert.Equal(t, "Test User", auth.FullName, "Expected FullName to be Test User")
	assert.Equal(t, "1234567890", auth.Phone, "Expected Phone to be 1234567890")
	assert.Equal(t, false, auth.IsSuperuser, "Expected IsSuperuser to be false")
}

func TestCustomClaims(t *testing.T) {
	claims := CustomClaims{
		UserID:      1,
		Username:    "testuser",
		Email:       "test@example.com",
		IsSuperuser: false,
	}

	assert.Equal(t, 1, claims.UserID, "UserID should be 1")
	assert.Equal(t, "testuser", claims.Username, "Username should be 'testuser'")
	assert.Equal(t, "test@example.com", claims.Email, "Email should be 'test@example.com'")
	assert.Equal(t, false, claims.IsSuperuser, "IsSuperuser should be false")
}
