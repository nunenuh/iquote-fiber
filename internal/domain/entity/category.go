package entity

import "time"

type Category struct {
	ID          int
	Name        string
	Description string
	ParentID    int
	Child       []Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	IsDeleted   bool
}
