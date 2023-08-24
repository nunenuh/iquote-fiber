package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	login := Login{
		Username: "testUser",
		Password: "testPassword",
	}

	// Use testify/assert for assertions
	assert.Equal(t, "testUser", login.Username)
	assert.Equal(t, "testPassword", login.Password)

	// Add more assertions to test the behavior of the struct
	assert.NotEmpty(t, login.Username, "Expected username to be non-empty")
	assert.NotEmpty(t, login.Password, "Expected password to be non-empty")
}
