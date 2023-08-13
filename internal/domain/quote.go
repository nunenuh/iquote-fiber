package domain

import "time"

type Quote struct {
	ID           int
	text         string
	tags         string
	Author       Author
	Category     []Category
	UserWhoLiked []User
	LikedCount   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	IsDeleted    bool
}
