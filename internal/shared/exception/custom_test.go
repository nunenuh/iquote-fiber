package exception

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepositoryError(t *testing.T) {
	msg := "Repository error"
	err := NewRepositoryError(msg)

	assert.Equal(t, RepositoryError, err.Type, "Expected error type to be RepositoryError")
	assert.Equal(t, msg, err.Message, "Expected error message to be %q", msg)
}

func TestNewValidatorError(t *testing.T) {
	msg := "Validator error"
	err := NewValidatorError(msg)

	assert.Equal(t, ValidatorError, err.Type, "Expected error type to be ValidatorError")
	assert.Equal(t, msg, err.Message, "Expected error message to be %q", msg)
}

func TestNewOtherError(t *testing.T) {
	msg := "Other error"
	err := NewOtherError(msg)

	assert.Equal(t, OtherError, err.Type, "Expected error type to be OtherError")
	assert.Equal(t, msg, err.Message, "Expected error message to be %q", msg)
}
