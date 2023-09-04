package infra

import (
	"testing"

	"github.com/nunenuh/iquote-fiber/internal/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/database/model"
	"github.com/stretchr/testify/assert"
)

func TestUserMapper_ToEntity(t *testing.T) {
	mapper := NewUserMapper()

	modelUser := &model.User{
		ID:       1,
		Email:    "test@example.com",
		Username: "testuser",
	}

	expectedAuth := &domain.Auth{
		ID:       1,
		Email:    "test@example.com",
		Username: "testuser",
	}

	auth := mapper.ToEntity(modelUser)

	assert.Equal(t, expectedAuth, auth)
}

func TestUserMapper_ToEntityWithPassword(t *testing.T) {
	mapper := NewUserMapper()

	modelUser := &model.User{
		ID:       1,
		Email:    "test@example.com",
		Username: "testuser",
		Password: "password123",
	}

	expectedAuth := &domain.Auth{
		ID:       1,
		Email:    "test@example.com",
		Password: "password123",
		Username: "testuser",
	}

	auth := mapper.ToEntityWithPassword(modelUser)

	assert.Equal(t, expectedAuth, auth)
}

func TestUserMapper_ToEntityList(t *testing.T) {
	mapper := NewUserMapper()

	modelUsers := []model.User{
		{
			ID:       1,
			Email:    "test1@example.com",
			Username: "testuser1",
		},
		{
			ID:       2,
			Email:    "test2@example.com",
			Username: "testuser2",
		},
	}

	expectedAuths := []*domain.Auth{
		{
			ID:       1,
			Email:    "test1@example.com",
			Username: "testuser1",
		},
		{
			ID:       2,
			Email:    "test2@example.com",
			Username: "testuser2",
		},
	}

	auths := mapper.ToEntityList(modelUsers)

	assert.Equal(t, expectedAuths, auths)
}

func TestUserMapper_ToModel(t *testing.T) {
	mapper := NewUserMapper()

	domainAuth := &domain.Auth{
		ID:       1,
		Email:    "test@example.com",
		Username: "testuser",
		Password: "password123",
	}

	expectedModel := &model.User{
		ID:       1,
		Email:    "test@example.com",
		Username: "testuser",
		Password: "password123",
		IsActive: true,
	}

	model := mapper.ToModel(domainAuth)

	assert.Equal(t, expectedModel, model)
}
