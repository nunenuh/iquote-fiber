package domain

import (
	"time"

	author "github.com/nunenuh/iquote-fiber/internal/core/author/domain"
	category "github.com/nunenuh/iquote-fiber/internal/core/category/domain"
	user "github.com/nunenuh/iquote-fiber/internal/core/user/domain"
)

type Quote struct {
	ID           int    `json:"id,omitempty"`
	QText        string `json:"qtext"`
	Tags         string `json:"tags,omitempty"`
	Author       author.Author
	Category     []category.Category
	UserWhoLiked []user.User
	LikedCount   int
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
	IsDeleted    bool      `json:"is_deleted,omitempty"`
}
