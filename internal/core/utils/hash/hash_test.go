package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswordPositive(t *testing.T) {
	password := "password123"
	hash, err := HashPassword(password)

	assert.NoError(t, err, "failed to hash password")
	assert.True(t, CheckHashPassword(password, hash), "CheckHashPassword returned false for a valid password")
}

func TestHashPasswordNegative(t *testing.T) {
	password := "password123"
	hash, err := HashPassword(password)

	assert.NoError(t, err, "failed to hash password")
	assert.False(t, CheckHashPassword("wrongpassword", hash), "CheckHashPassword returned true for an invalid password")
}

func TestCheckHashPassword(t *testing.T) {
	password := "password123"
	hash, _ := HashPassword(password)

	assert.True(t, CheckHashPassword(password, hash), "CheckHashPassword returned false for a valid password")

	invalidPassword := "wrongpassword"
	assert.False(t, CheckHashPassword(invalidPassword, hash), "CheckHashPassword returned true for an invalid password")
}
