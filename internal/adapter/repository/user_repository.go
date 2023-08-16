package repository

import (
	"fmt"
	"strconv"

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
		ID:       strconv.Itoa(user.ID),
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}
	return out, nil
}

func (r *userRepository) GetAll(limit int, offset int) ([]*entity.User, error) {
	db := r.DB
	var user []model.User
	result := db.Offset(offset).Limit(limit).Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	out := make([]*entity.User, 0)
	for _, u := range user {
		out = append(out, &entity.User{
			ID:       strconv.Itoa(u.ID),
			FullName: u.FullName,
			Email:    u.Email,
			Password: u.Password,
		})
	}
	return out, nil
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	db := r.DB

	userModel := &model.User{
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: true,
	}

	result := db.Create(&userModel)
	if result.Error != nil {
		panic(result.Error)
	}
	return user, nil
}

func (r *userRepository) Update(ID int, user *entity.User) (*entity.User, error) {
	db := r.DB

	userModel := &model.User{
		ID:       ID,
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}

	result := db.Save(&userModel)
	if result.Error != nil {
		panic(result.Error)
	}
	return user, nil
}

func (r *userRepository) Delete(ID int) error {
	db := r.DB

	var user model.User

	// Check if the user with the given ID exists
	if err := db.First(&user, ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("User with ID %d not found", ID)
		}
		return err
	}

	// Delete the user
	result := db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
