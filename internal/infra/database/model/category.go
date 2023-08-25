package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	ParentID    *int      `gorm:"index;default:NULL" json:"parent_id"`
	Parent      *Category `gorm:"foreignKey:ParentID"`
	Quotes      []Quote   `gorm:"many2many:quote_categories"`
}

func (q *Category) IsDeleted() bool {
	return !q.DeletedAt.Time.IsZero()
}

// Ensure foreign key constraints are applied on migrations.
func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if category.ParentID != nil && *category.ParentID == 0 {
		category.ParentID = nil
	}
	return
}
