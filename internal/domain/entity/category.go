package entity

import "time"

type Category struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name" validate:"required,min=2,max=100"`
	Description string     `json:"description,omitempty" validate:"required,min=100,max=500"`
	ParentID    string     `json:"parent_id,omitempty"`
	Child       []Category `json:"child,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	DeletedAt   time.Time  `json:"deleted_at,omitempty"`
	IsDeleted   bool       `json:"is_deleted,omitempty"`
}
