package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	FullName string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Phone    string `json:"phone"`
	Level    string `json:"level"`
	Active   bool   `json:"active"`
}
