package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthor(t *testing.T) {
	now := time.Now()
	author := Author{
		ID:        "1",
		Name:      "John Doe",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: now,
		IsDeleted: false,
	}

	// Add your test assertions here
	assert.Equal(t, "1", author.ID, "Expected ID to be 1")
	assert.Equal(t, "John Doe", author.Name, "Expected Name to be 'John Doe'")
	assert.True(t, author.CreatedAt.Equal(now) || author.CreatedAt.Before(now), "Expected CreatedAt to be equal or before current time")
	assert.True(t, author.UpdatedAt.Equal(now) || author.UpdatedAt.Before(now), "Expected UpdatedAt to be equal or before current time")
	assert.True(t, author.DeletedAt.Equal(now) || author.DeletedAt.Before(now), "Expected DeletedAt to be equal or before current time")
	assert.False(t, author.IsDeleted, "Expected IsDeleted to be false")
}
