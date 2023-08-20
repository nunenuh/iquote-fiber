package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	if !CheckHashPassword(password, hash) {
		t.Errorf("CheckHashPassword returned false for a valid password")
	}
}

func TestCheckHashPassword(t *testing.T) {
	password := "password123"
	hash, _ := HashPassword(password)

	if !CheckHashPassword(password, hash) {
		t.Errorf("CheckHashPassword returned false for a valid password")
	}

	invalidPassword := "wrongpassword"
	if CheckHashPassword(invalidPassword, hash) {
		t.Errorf("CheckHashPassword returned true for an invalid password")
	}
}
