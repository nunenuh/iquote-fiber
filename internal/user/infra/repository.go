package user

import (
	"fmt"
	"log"

	"github.com/nunenuh/iquote-fiber/internal/database/model"
	"github.com/nunenuh/iquote-fiber/internal/user/domain"
	"gorm.io/gorm"
)

func ProvideUserRepository(db *gorm.DB) domain.IUserRepository {
	return NewUserRepository(db)
}

type userRepository struct {
	DB     *gorm.DB
	Mapper *UserMapper
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB:     db,
		Mapper: NewUserMapper(),
	}
}

func (r *userRepository) GetAll(limit int, offset int) ([]*domain.User, error) {
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

func (r *userRepository) GetByID(ID int) (*domain.User, error) {
	user, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	out := r.Mapper.ToEntity(user)
	return out, nil
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	db := r.DB
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User with username %s not found", username)
		}
		return nil, result.Error
	}
	log.Printf("Password from repo: %s", user.Password)

	out := r.Mapper.ToEntityWithPassword(&user)
	return out, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
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

func (r *userRepository) Create(user *domain.User) (*domain.User, error) {
	db := r.DB

	userModel := r.Mapper.ToModel(user)

	result := db.Create(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	out := r.Mapper.ToEntity(userModel)
	return out, nil
}

func (r *userRepository) Update(ID int, user *domain.User) (*domain.User, error) {
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
