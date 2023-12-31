package infra

import (
	"fmt"

	"github.com/nunenuh/iquote-fiber/internal/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/database/model"
	"gorm.io/gorm"
)

func ProvideAuthRepository(db *gorm.DB) domain.IAuthRepository {
	return NeAuthRepository(db)
}

type authRepository struct {
	DB     *gorm.DB
	Mapper *UserMapper
}

func NeAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		DB:     db,
		Mapper: NewUserMapper(),
	}
}

func (r *authRepository) FindByID(ID int) (*model.User, error) {
	db := r.DB
	var user model.User
	result := db.First(&user, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("User with ID %d not found!", ID)
	}

	return &user, nil
}

func (r *authRepository) GetByUsername(username string) (*domain.Auth, error) {
	db := r.DB
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User with username %s not found", username)
		}
		return nil, result.Error
	}

	out := r.Mapper.ToEntityWithPassword(&user)
	return out, nil
}
