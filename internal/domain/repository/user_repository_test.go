package repository

import (
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct {
	GetAllFunc        func(limit int, offset int) ([]*entity.User, error)
	GetByIDFunc       func(ID int) (*entity.User, error)
	GetByUsernameFunc func(username string) (*entity.User, error)
	GetByEmailFunc    func(email string) (*entity.User, error)
	CreateFunc        func(user *entity.User) (*entity.User, error)
	UpdateFunc        func(ID int, user *entity.User) (*entity.User, error)
	DeleteFunc        func(ID int) error
}

func (m *mockUserRepository) GetAll(limit int, offset int) ([]*entity.User, error) {
	return m.GetAllFunc(limit, offset)
}

func (m *mockUserRepository) GetByID(ID int) (*entity.User, error) {
	return m.GetByIDFunc(ID)
}

func (m *mockUserRepository) GetByUsername(username string) (*entity.User, error) {
	return m.GetByUsernameFunc(username)
}

func (m *mockUserRepository) GetByEmail(email string) (*entity.User, error) {
	return m.GetByEmailFunc(email)
}

func (m *mockUserRepository) Create(user *entity.User) (*entity.User, error) {
	return m.CreateFunc(user)
}

func (m *mockUserRepository) Update(ID int, user *entity.User) (*entity.User, error) {
	return m.UpdateFunc(ID, user)
}

func (m *mockUserRepository) Delete(ID int) error {
	return m.DeleteFunc(ID)
}
func TestMockUserRepositoryGetAll(t *testing.T) {
	// Create an instance of the mockUserRepository
	mockRepo := &mockUserRepository{
		GetAllFunc: func(limit int, offset int) ([]*entity.User, error) {
			// Implement the mock logic for GetAllFunc
			users := []*entity.User{
				{ID: 1, FullName: "John Doe"},
				{ID: 2, FullName: "Jane Smith"},
			}
			return users, nil
		},
	}

	// Call the GetAll method on the mock repository
	users, err := mockRepo.GetAll(10, 0)

	// Assert that the returned users are as expected
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "John Doe", users[0].FullName)
	assert.Equal(t, "Jane Smith", users[1].FullName)
}

func TestMockUserRepositoryGetByID(t *testing.T) {
	// Create an instance of the mockUserRepository
	mockRepo := &mockUserRepository{
		GetByIDFunc: func(id int) (*entity.User, error) {
			// Implement the mock logic for GetByIDFunc
			user := &entity.User{
				ID:       id,
				FullName: "John Doe",
			}
			return user, nil
		},
	}

	// Call the GetByID method on the mock repository
	user, err := mockRepo.GetByID(1)

	// Assert that the returned user is as expected
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.FullName)
}

func TestMockUserRepositoryGetByUsername(t *testing.T) {
	// Create an instance of the mockUserRepository
	mockRepo := &mockUserRepository{
		GetByUsernameFunc: func(username string) (*entity.User, error) {
			// Implement the mock logic for GetByUsernameFunc
			user := &entity.User{
				ID:       1,
				FullName: "John Doe",
				Username: username,
			}
			return user, nil
		},
	}

	// Call the GetByUsername method on the mock repository
	user, err := mockRepo.GetByUsername("johndoe")

	// Assert that the returned user is as expected
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.FullName)
	assert.Equal(t, "johndoe", user.Username)
}

func TestMockUserRepositoryGetByEmail(t *testing.T) {
	// Create an instance of the mockUserRepository
	mockRepo := &mockUserRepository{
		GetByEmailFunc: func(email string) (*entity.User, error) {
			// Implement the mock logic for GetByEmailFunc
			user := &entity.User{
				ID:       1,
				FullName: "John Doe",
				Email:    email,
			}
			return user, nil
		},
	}

	// Call the GetByEmail method on the mock repository
	user, err := mockRepo.GetByEmail("johndoe@example.com")

	// Assert that the returned user is as expected
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.FullName)
	assert.Equal(t, "johndoe@example.com", user.Email)
}

func TestMockUserRepositoryCreate(t *testing.T) {
	// Create an instance of the mockUserRepository
	mockRepo := &mockUserRepository{
		CreateFunc: func(user *entity.User) (*entity.User, error) {
			// Implement the mock logic for CreateFunc
			user.ID = 1
			return user, nil
		},
	}

	// Create a user entity
	user := &entity.User{
		FullName: "John Doe",
	}

	// Call the Create method on the mock repository
	createdUser, err := mockRepo.Create(user)

	// Assert that the created user is as expected
	assert.NoError(t, err)
	assert.Equal(t, 1, createdUser.ID)
	assert.Equal(t, "John Doe", createdUser.FullName)
}

func TestMockUserRepositoryUpdate(t *testing.T) {
	// Create an instance of the mockUserRepository
	mockRepo := &mockUserRepository{
		UpdateFunc: func(ID int, user *entity.User) (*entity.User, error) {
			// Implement the mock logic for UpdateFunc
			user.ID = ID
			return user, nil
		},
	}

	// Create a user entity
	user := &entity.User{
		ID:       1,
		FullName: "John Doe",
	}

	// Call the Update method on the mock repository
	updatedUser, err := mockRepo.Update(1, user)

	// Assert that the updated user is as expected
	assert.NoError(t, err)
	assert.Equal(t, 1, updatedUser.ID)
	assert.Equal(t, "John Doe", updatedUser.FullName)
}

func TestMockUserRepositoryDelete(t *testing.T) {
	// Create an instance of the mockUserRepository
	mockRepo := &mockUserRepository{
		DeleteFunc: func(ID int) error {
			// Implement the mock logic for DeleteFunc
			return nil
		},
	}

	// Call the Delete method on the mock repository
	err := mockRepo.Delete(1)

	// Assert that the error is as expected
	assert.NoError(t, err)

}
