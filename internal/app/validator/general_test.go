package validator

import (
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestValidator_Validate(t *testing.T) {
	v := NewValidator()

	// Test case 1: Valid entity
	entity := generateValidEntity()
	assert.NoError(t, v.Validate(entity))

	// Test case 2: Invalid entity
	invalidEntity := generateInvalidEntity()
	assert.Error(t, v.Validate(invalidEntity))
}

func generateValidEntity() entity.User {
	return entity.User{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
		IsActive: true,
	}
}

func generateInvalidEntity() entity.User {
	// create and return an invalid entity
	return entity.User{
		FullName: "John Doe",
		Email:    "johndoe-wrong-example.com",
		IsActive: true,
	}
}
