package domain

import (
	"testing"
	"time"
)

func TestAuthor(t *testing.T) {
	now := time.Now()
	author := Author{
		ID:        1,
		Name:      "John Doe",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: now,
		IsDeleted: false,
	}

	// Add your test assertions here
	if author.ID != 1 {
		t.Errorf("Expected ID to be 1, but got %d", author.ID)
	}

	if author.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', but got '%s'", author.Name)
	}

	if !author.CreatedAt.Equal(now) && !author.CreatedAt.Before(now) {
		t.Errorf("Expected CreatedAt to be equal or before current time, but got '%s'", author.CreatedAt)
	}

	if !author.UpdatedAt.Equal(now) && !author.UpdatedAt.Before(now) {
		t.Errorf("Expected UpdatedAt to be equal or before current time, but got '%s'", author.UpdatedAt)
	}

	if !author.DeletedAt.Equal(now) && !author.DeletedAt.Before(now) {
		t.Errorf("Expected DeletedAt to be equal or before current time, but got '%s'", author.DeletedAt)
	}

	if author.IsDeleted {
		t.Errorf("Expected IsDeleted to be false, but got true")
	}
}
