package domain

import (
	"testing"
	"time"

	author "github.com/nunenuh/iquote-fiber/internal/author/domain"
	category "github.com/nunenuh/iquote-fiber/internal/category/domain"
	user "github.com/nunenuh/iquote-fiber/internal/user/domain"
	"github.com/stretchr/testify/assert"
)

func TestQuote(t *testing.T) {
	quote := Quote{
		ID:           1,
		QText:        "Test quote",
		Tags:         "test",
		Author:       author.Author{},
		Category:     []category.Category{},
		UserWhoLiked: []user.User{},
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
	assert.Equal(t, author.Author{}, quote.Author)
	assert.Equal(t, []category.Category{}, quote.Category)
	assert.Equal(t, []user.User{}, quote.UserWhoLiked)
	assert.Equal(t, 0, quote.LikedCount)
	assert.Equal(t, false, quote.IsDeleted)
	assert.WithinDuration(t, time.Now(), quote.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), quote.UpdatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), quote.DeletedAt, time.Second)

	// Add more assertions as needed
}
