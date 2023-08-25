package author

import (
	"fmt"

	"github.com/nunenuh/iquote-fiber/internal/core/author/domain"
	"github.com/nunenuh/iquote-fiber/internal/infra/database/model"
	"gorm.io/gorm"
)

func ProvideAuthorRepository(db *gorm.DB) domain.IAuthorRepository {
	return NewAuthorRepository(db)
}

type authorRepository struct {
	DB     *gorm.DB
	Mapper *AuthorMapper
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{
		DB:     db,
		Mapper: NewAuthorMapper(),
	}
}

func (r *authorRepository) GetAll(limit int, offset int) ([]*domain.Author, error) {
	db := r.DB
	var authorModel []model.Author
	result := db.Offset(offset).Limit(limit).Find(&authorModel)
	if result.Error != nil {
		return nil, result.Error
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

func (r *authorRepository) GetByID(ID int) (*domain.Author, error) {
	authorModel, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}
	out := r.Mapper.ToEntity(authorModel)
	return out, nil
}

func (r *authorRepository) GetByName(name string) (*domain.Author, error) {
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

func (r *authorRepository) Create(author *domain.Author) (*domain.Author, error) {
	db := r.DB
	authorModel := r.Mapper.ToModel(author)
	result := db.Create(&authorModel)
	if result.Error != nil {
		return nil, result.Error
	}
	author = r.Mapper.ToEntity(authorModel)
	return author, nil
}

func (r *authorRepository) Update(ID int, author *domain.Author) (*domain.Author, error) {
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
