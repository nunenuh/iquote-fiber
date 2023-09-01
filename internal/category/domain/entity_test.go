package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {
	now := time.Now()

	category := Category{
		ID:          1,
		Name:        "Category 1",
		Description: "Description of Category 1",
		ParentID:    0,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   now,
		IsDeleted:   false,
	}

	assert.Equal(t, 1, category.ID, "Expected ID to be 1")
	assert.Equal(t, "Category 1", category.Name, "Expected Name to be 'Category 1'")
	assert.Equal(t, "Description of Category 1", category.Description, "Expected Description to be 'Description of Category 1'")
	assert.Equal(t, 0, category.ParentID, "Expected ParentID to be 0")
	assert.Equal(t, now, category.CreatedAt, "Expected CreatedAt to be equal to current time")
	assert.Equal(t, now, category.UpdatedAt, "Expected UpdatedAt to be equal to current time")
	assert.Equal(t, now, category.DeletedAt, "Expected DeletedAt to be equal to current time")
	assert.False(t, category.IsDeleted, "Expected IsDeleted to be false")
}
