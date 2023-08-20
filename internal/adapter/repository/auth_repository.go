package repository

import (
	"fmt"
	"strconv"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideAuthRepository(db *gorm.DB) repository.IAuthRepository {
	return NewAuthRepository(db)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		DB: db,
	}
}

func (r *authRepository) Login(username string, password string) (*entity.User, error) {
	db := r.DB
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User with username %s not found", username)
		}
		return nil, result.Error
	}

	out := &entity.User{
		ID:          strconv.Itoa(user.ID),
		FullName:    user.FullName,
		Email:       user.Email,
		Password:    user.Password,
		IsActive:    user.IsActive,
		Username:    user.Username,
		Phone:       user.Phone,
		Level:       user.Level,
		IsSuperuser: user.IsSuperuser,
	}
	return out, nil
}
