package domain

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	now := time.Now()
	user := User{
		ID:          "1",
		Username:    "john_doe",
		Password:    "password123",
		FullName:    "John Doe",
		Email:       "john.doe@example.com",
		Phone:       "1234567890",
		IsActive:    true,
		IsSuperuser: false,
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   now,
		IsDeleted:   false,
	}

	if user.ID != "1" {
		t.Errorf("Expected ID = %v, but got %v", "1", user.ID)
	}

	if user.Username != "john_doe" {
		t.Errorf("Expected Username = %v, but got %v", "john_doe", user.Username)
	}

	if user.Password != "password123" {
		t.Errorf("Expected Password = %v, but got %v", "password123", user.Password)
	}

	if user.FullName != "John Doe" {
		t.Errorf("Expected Name = %v, but got %v", "John Doe", user.FullName)
	}

	if user.Email != "john.doe@example.com" {
		t.Errorf("Expected Email = %v, but got %v", "john.doe@example.com", user.Email)
	}

	if user.Phone != "1234567890" {
		t.Errorf("Expected Phone = %v, but got %v", "1234567890", user.Phone)
	}

	if !user.IsActive {
		t.Errorf("Expected IsActive = %v, but got %v", true, user.IsActive)
	}

	if user.IsSuperuser {
		t.Errorf("Expected IsSuperuser = %v, but got %v", false, user.IsSuperuser)
	}

	if user.CreatedAt != now {
		t.Errorf("Expected CreatedAt = %v, but got %v", now, user.CreatedAt)
	}

	if user.UpdatedAt != now {
		t.Errorf("Expected UpdatedAt = %v, but got %v", now, user.UpdatedAt)
	}

	if user.DeletedAt != now {
		t.Errorf("Expected DeletedAt = %v, but got %v", now, user.DeletedAt)
	}

	if user.IsDeleted {
		t.Errorf("Expected IsDeleted = %v, but got %v", false, user.IsDeleted)
	}
}
