package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQuote(t *testing.T) {
	quote := Quote{
		ID:           1,
		QText:        "Test quote",
		Tags:         "test",
		Author:       Author{},
		Category:     []Category{},
		UserWhoLiked: []User{},
		LikedCount:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    time.Now(),
		IsDeleted:    false,
	}

	// Add your test cases here
	assert.Equal(t, 1, quote.ID)
	assert.Equal(t, "Test quote", quote.QText)
	assert.Equal(t, "test", quote.Tags)
	assert.Equal(t, Author{}, quote.Author)
	assert.Equal(t, []Category{}, quote.Category)
	assert.Equal(t, []User{}, quote.UserWhoLiked)
	assert.Equal(t, 0, quote.LikedCount)
	assert.Equal(t, false, quote.IsDeleted)
	assert.WithinDuration(t, time.Now(), quote.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), quote.UpdatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), quote.DeletedAt, time.Second)

	// Add more assertions as needed
}
