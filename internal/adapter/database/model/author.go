package model

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model

	ID          int            `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique" json:"name"`
	Description string         `json:"description"`
	IsActive    bool           `json:"is_active"`
	IsSuperuser bool           `json:"is_superuser"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (a *Author) IsDeleted() bool {
	return !a.DeletedAt.Time.IsZero()
}
