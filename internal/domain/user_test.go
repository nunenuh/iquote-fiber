package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := User{
		ID:       "1",
		Username: "john_doe",
		Password: "password123",
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Phone:    "1234567890",
		Active:   true,
	}

	// Add your test assertions here
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "john_doe", user.Username)
	assert.Equal(t, "password123", user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "1234567890", user.Phone)
	assert.True(t, user.Active)
}
