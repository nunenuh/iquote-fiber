package repository

import "testing"

func TestGet(t *testing.T) {
	// Create a mock implementation of the IUserRepository interface
	repo := &UserRepository{} // Replace UserRepository with the actual implementation of IUserRepository

	// Call the Get method and assert the expected result
	user := repo.Get()
	// Add your assertions here
}
