package entity

import "time"

type User struct {
	ID          string    `json:"id,omitempty"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	FullName    string    `json:"full_name" validate:"required,min=2,max=100"`
	Email       string    `json:"email" validate:"required,email"`
	Phone       string    `json:"phone"`
	Level       string    `json:"level,omitempty"`
	IsActive    bool      `json:"is_active,omitempty"`
	IsSuperuser bool      `json:"is_superuser,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	IsDeleted   bool      `json:"is_deleted,omitempty"`
}
