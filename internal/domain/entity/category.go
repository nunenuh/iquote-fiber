package entity

import "time"

type Category struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name" validate:"required,min=2,max=100"`
	Description string     `json:"description,omitempty" validate:"required,min=10,max=500"`
	ParentID    int        `json:"parent_id,omitempty"`
	Parent      *Category  `json:"parent,omitempty"`
	Child       []Category `json:"child,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	DeletedAt   time.Time  `json:"deleted_at,omitempty"`
	IsDeleted   bool       `json:"is_deleted,omitempty"`
}
