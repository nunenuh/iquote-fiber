package repository

import (
	"fmt"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/adapter/mapper"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideUserRepository(db *gorm.DB) repository.IUserRepository {
	return NewUserRepository(db)
}

type userRepository struct {
	DB     *gorm.DB
	Mapper *mapper.UserMapper
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB:     db,
		Mapper: mapper.NewUserMapper(),
	}
}

func (r *userRepository) GetAll(limit int, offset int) ([]*entity.User, error) {
	db := r.DB
	var users []model.User
	result := db.Offset(offset).Limit(limit).Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("User data is empty!")
	}

	out := r.Mapper.ToEntityList(users)
	return out, nil
}

func (r *userRepository) FindByID(ID int) (*model.User, error) {
	db := r.DB
	var user model.User
	result := db.First(&user, ID)
	if result.Error != nil {
		return nil, fmt.Errorf("User with ID %d not found!", ID)
	}

	return &user, nil
}

func (r *userRepository) GetByID(ID int) (*entity.User, error) {
	user, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	out := r.Mapper.ToEntity(user)
	return out, nil
}

func (r *userRepository) GetByUsername(username string) (*entity.User, error) {
	db := r.DB
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User with username %s not found", username)
		}
		return nil, result.Error
	}

	out := r.Mapper.ToEntity(&user)
	return out, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	db := r.DB
	var user model.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User with email %s not found", email)
		}
		return nil, result.Error
	}

	out := r.Mapper.ToEntity(&user)
	return out, nil
}

func (r *userRepository) Create(user *entity.User) (*entity.User, error) {
	db := r.DB

	userModel := r.Mapper.ToModel(user)

	result := db.Create(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	out := r.Mapper.ToEntity(userModel)
	return out, nil
}

func (r *userRepository) Update(ID int, user *entity.User) (*entity.User, error) {
	db := r.DB

	userModel, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	user.ID = ID
	userModel = r.Mapper.ToModel(user)

	result := db.Save(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	out := r.Mapper.ToEntity(userModel)
	return out, nil
}

func (r *userRepository) Delete(ID int) error {
	db := r.DB

	user, err := r.FindByID(ID)
	if err != nil {
		return err
	}

	// Delete the user
	result := db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
