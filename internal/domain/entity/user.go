package entity

import "time"

type User struct {
	ID          int
	Username    string
	Password    string
	FullName    string
	Email       string
	Phone       string
	Level       string
	IsActive    bool
	IsSuperuser bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	IsDeleted   bool
}
