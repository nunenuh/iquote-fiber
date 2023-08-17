package entity

import "time"

type Quote struct {
	ID           string
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
