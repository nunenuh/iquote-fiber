package repository

import (
	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetByID(ID int) (*entity.User, error) {
	db := r.DB
	var user model.User
	result := db.First(&user, ID)
	if result.Error != nil {
		panic(result.Error)
	}

	out := &entity.User{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}
	return out, nil

}
