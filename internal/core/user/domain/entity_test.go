package domain

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// func TestUser(t *testing.T) {
// 	now := time.Now()
// 	user := User{
// 		ID:          1,
// 		Username:    "john_doe",
// 		Password:    "password123",
// 		FullName:    "John Doe",
// 		Email:       "john.doe@example.com",
// 		Phone:       "1234567890",
// 		IsActive:    true,
// 		IsSuperuser: false,
// 		CreatedAt:   now,
// 		UpdatedAt:   now,
// 		DeletedAt:   now,
// 		IsDeleted:   false,
// 	}

// 	assert.Equal(t, 1, user.ID, "Expected ID = %v, but got %v")
// 	assert.Equal(t, "john_doe", user.Username, "Expected Username = %v, but got %v")
// 	assert.Equal(t, "password123", user.Password, "Expected Password = %v, but got %v")
// 	assert.Equal(t, "John Doe", user.FullName, "Expected Name = %v, but got %v")
// 	assert.Equal(t, "john.doe@example.com", user.Email, "Expected Email = %v, but got %v")
// 	assert.Equal(t, "1234567890", user.Phone, "Expected Phone = %v, but got %v")
// 	assert.Equal(t, true, user.IsActive, "Expected IsActive = %v, but got %v")
// 	assert.Equal(t, false, user.IsSuperuser, "Expected IsSuperuser = %v, but got %v")
// 	assert.Equal(t, now, user.CreatedAt, "Expected CreatedAt = %v, but got %v")
// 	assert.Equal(t, now, user.UpdatedAt, "Expected UpdatedAt = %v, but got %v")
// 	assert.Equal(t, now, user.DeletedAt, "Expected DeletedAt = %v, but got %v")
// 	assert.Equal(t, false, user.IsDeleted, "Expected IsDeleted = %v, but got %v")
// }
