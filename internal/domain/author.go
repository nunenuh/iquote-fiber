package domain

import "time"

type Author struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	IsDeleted   bool
}
