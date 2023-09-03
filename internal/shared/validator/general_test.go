package validator

import (
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/user/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidatorValidate(t *testing.T) {
	v := NewValidator()

	// Test case 1: Valid entity
	entity := generateValidEntity()
	assert.NoError(t, v.Validate(entity))

	// Test case 2: Invalid entity
	invalidEntity := generateInvalidEntity()
	assert.Error(t, v.Validate(invalidEntity))
}

func generateValidEntity() domain.User {
	return domain.User{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		IsActive: true,
	}
}

func generateInvalidEntity() domain.User {
	// create and return an invalid entity
	return domain.User{
		FullName: "John Doe",
		Email:    "johndoe-wrong-example.com",
		IsActive: true,
	}
}
