package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:255" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	ParentID    *int           `gorm:"index;default:NULL" json:"parent_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (u *Category) IsDeleted() bool {
	return !u.DeletedAt.Time.IsZero()
}

// Ensure foreign key constraints are applied on migrations.
func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if category.ParentID != nil && *category.ParentID == 0 {
		category.ParentID = nil
	}
	return
}
