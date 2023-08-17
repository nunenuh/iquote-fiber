package entity

import "time"

type Category struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	ParentID    string     `json:"parent_id,omitempty"`
	Child       []Category `json:"child,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	DeletedAt   time.Time  `json:"deleted_at,omitempty"`
	IsDeleted   bool       `json:"is_deleted,omitempty"`
}
