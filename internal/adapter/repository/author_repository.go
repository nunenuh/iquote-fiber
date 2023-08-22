package repository

import (
	"fmt"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideAuthorRepository(db *gorm.DB) repository.IAuthorRepository {
	return NewAuthorRepository(db)
}

type authorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{
		DB: db,
	}
}

func (r *authorRepository) GetAll(limit int, offset int) ([]*entity.Author, error) {
	db := r.DB
	var authorModel []model.Author
	result := db.Offset(offset).Limit(limit).Find(&authorModel)
	if result.Error != nil {
		panic(result.Error)
	}

	out := make([]*entity.Author, 0)
	for _, u := range authorModel {
		out = append(out, &entity.Author{
			ID:          u.ID,
			Name:        u.Name,
			Description: u.Description,
		})
	}
	return out, nil
}

func (r *authorRepository) GetByID(ID int) (*entity.Author, error) {
	db := r.DB
	var authorModel model.Author
	result := db.First(&authorModel, ID)
	if result.Error != nil {
		panic(result.Error)
	}

	out := &entity.Author{
		ID:          authorModel.ID,
		Name:        authorModel.Name,
		Description: authorModel.Description,
	}
	return out, nil
}

func (r *authorRepository) GetByName(name string) (*entity.Author, error) {
	db := r.DB
	var authorModel model.Author
	result := db.Where("name = ?", name).First(&authorModel)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User with name %s not found", name)
		}
		return nil, result.Error
	}

	out := &entity.Author{
		ID:          authorModel.ID,
		Name:        authorModel.Name,
		Description: authorModel.Description,
	}
	return out, nil
}

func (r *authorRepository) Create(author *entity.Author) (*entity.Author, error) {
	db := r.DB

	authorModel := &model.Author{
		Name:        author.Name,
		Description: author.Description,
		IsActive:    true,
	}

	result := db.Create(&authorModel)
	if result.Error != nil {
		panic(result.Error)
	}
	return author, nil
}

func (r *authorRepository) Update(ID int, author *entity.Author) (*entity.Author, error) {
	db := r.DB

	authorModel := &model.Author{
		ID:          ID,
		Name:        author.Name,
		Description: author.Description,
		IsActive:    true,
	}

	result := db.Save(&authorModel)
	if result.Error != nil {
		panic(result.Error)
	}
	return author, nil
}

func (r *authorRepository) Delete(ID int) error {
	db := r.DB

	var authorModel model.Author

	// Check if the author with the given ID exists
	if err := db.First(&authorModel, ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("author with ID %d not found", ID)
		}
		return err
	}

	// Delete the author
	result := db.Delete(&authorModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
