package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID          int            `gorm:"primaryKey" json:"id"`
	Username    string         `gorm:"unique" json:"username"`
	Password    string         `json:"-"`
	FullName    string         `json:"name"`
	Email       string         `gorm:"unique" json:"email"`
	Phone       string         `json:"phone"`
	Level       string         `json:"level"`
	IsActive    bool           `json:"is_active"`
	IsSuperuser bool           `json:"is_superuser"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) IsDeleted() bool {
	return !u.DeletedAt.Time.IsZero()
}
