package repository

import (
	"fmt"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/adapter/mapper"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideAuthorRepository(db *gorm.DB) repository.IAuthorRepository {
	return NewAuthorRepository(db)
}

type authorRepository struct {
	DB     *gorm.DB
	Mapper *mapper.AuthorMapper
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{
		DB:     db,
		Mapper: mapper.NewAuthorMapper(),
	}
}

func (r *authorRepository) GetAll(limit int, offset int) ([]*entity.Author, error) {
	db := r.DB
	var authorModel []model.Author
	result := db.Offset(offset).Limit(limit).Find(&authorModel)
	if result.Error != nil {
		panic(result.Error)
	}

	out := r.Mapper.ToEntityList(authorModel)
	return out, nil
}

func (r *authorRepository) FindByID(ID int) (*model.Author, error) {
	db := r.DB
	var authorModel model.Author
	result := db.First(&authorModel, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &authorModel, nil
}

func (r *authorRepository) GetByID(ID int) (*entity.Author, error) {
	authorModel, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}
	out := r.Mapper.ToEntity(authorModel)
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

	out := r.Mapper.ToEntity(&authorModel)
	return out, nil
}

func (r *authorRepository) Create(author *entity.Author) (*entity.Author, error) {
	db := r.DB
	authorModel := r.Mapper.ToModel(author)
	result := db.Create(&authorModel)
	if result.Error != nil {
		return nil, result.Error
	}
	author = r.Mapper.ToEntity(authorModel)
	return author, nil
}

func (r *authorRepository) Update(ID int, author *entity.Author) (*entity.Author, error) {
	db := r.DB

	authorModel, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	author.ID = ID
	authorModel = r.Mapper.ToModel(author)

	result := db.Save(&authorModel)
	if result.Error != nil {
		return nil, result.Error
	}
	author = r.Mapper.ToEntity(authorModel)
	return author, nil
}

func (r *authorRepository) Delete(ID int) error {
	db := r.DB

	authorModel, err := r.FindByID(ID)
	if err != nil {
		return err
	}

	// Delete the author
	result := db.Delete(&authorModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
