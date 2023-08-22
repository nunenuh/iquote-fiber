package entity

import "time"

type Quote struct {
	ID           int    `json:"id,omitempty"`
	QText        string `json:"qtext"`
	Tags         string `json:"tags,omitempty"`
	Author       Author
	Category     []Category
	UserWhoLiked []User
	LikedCount   int
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
	IsDeleted    bool      `json:"is_deleted,omitempty"`
}
