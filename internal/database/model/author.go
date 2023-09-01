package model

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model

	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	IsSuperuser bool   `json:"is_superuser"`
}

func (a *Author) IsDeleted() bool {
	return !a.DeletedAt.Time.IsZero()
}
