package repository

import (
	"fmt"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/adapter/mapper"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideCategoryRepository(db *gorm.DB) repository.ICategoryRepository {
	return NewCategoryRepository(db)
}

type categoryRepository struct {
	DB     *gorm.DB
	Mapper mapper.CategoryMapper
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		DB:     db,
		Mapper: *mapper.NewCategoryMapper(),
	}
}

func (r *categoryRepository) GetAll(limit int, offset int) ([]*entity.Category, error) {
	db := r.DB
	var categoryModel []model.Category
	result := db.Preload("Parent").Offset(offset).Limit(limit).Find(&categoryModel)
	if result.Error != nil {
		return nil, result.Error
	}

	out := r.Mapper.ToEntityList(categoryModel)
	return out, nil
}

func (r *categoryRepository) FindByID(ID int) (*model.Category, error) {
	db := r.DB
	var categoryModel model.Category
	result := db.First(&categoryModel, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &categoryModel, nil

}

func (r *categoryRepository) GetByID(ID int) (*entity.Category, error) {
	categoryModel, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	out := r.Mapper.ToEntity(categoryModel)
	return out, nil
}

func (r *categoryRepository) GetByName(name string) ([]*entity.Category, error) {
	db := r.DB
	var categoryModel []model.Category
	result := db.Where("name = ?", name).Find(&categoryModel)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Category with name %s not found", name)
		}
		return nil, result.Error
	}
	out := r.Mapper.ToEntityList(categoryModel)
	return out, nil
}

func (r *categoryRepository) FindByParentID(ID int) ([]model.Category, error) {
	db := r.DB
	var categoryModel []model.Category
	result := db.Where("parent_id = ?", ID).Find(&categoryModel)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Category with parent_id %d not found", ID)
		}
		return nil, result.Error
	}
	return categoryModel, nil
}

func (r *categoryRepository) GetByParentID(ID int) ([]*entity.Category, error) {
	db := r.DB
	var categoryModel []model.Category
	result := db.Where("parent_id = ?", ID).Find(&categoryModel)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Category with parent_id %d not found", ID)
		}
		return nil, result.Error
	}

	out := r.Mapper.ToEntityList(categoryModel)
	return out, nil
}

func (r *categoryRepository) Create(category *entity.Category) (*entity.Category, error) {
	db := r.DB
	categoryModel := &model.Category{
		Name:        category.Name,
		Description: category.Description,
	}

	if category.ParentID != 0 {
		parentCategory, err := r.FindByID(category.ParentID)
		if err != nil {
			return nil, err
		}
		categoryModel.ParentID = &parentCategory.ID
	}

	result := db.Create(&categoryModel)
	if result.Error != nil {
		return nil, result.Error
	}

	category.CreatedAt = categoryModel.CreatedAt
	category.UpdatedAt = categoryModel.UpdatedAt
	out := r.Mapper.ToEntity(categoryModel)

	return out, nil
}

func (r *categoryRepository) Update(ID int, category *entity.Category) (*entity.Category, error) {
	db := r.DB

	categoryModel, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}
	categoryModel.Name = category.Name
	categoryModel.Description = category.Description

	if category.ParentID != 0 {
		parentCategory, err := r.FindByID(category.ParentID)
		if err != nil {
			return nil, err
		}
		categoryModel.ParentID = &parentCategory.ID
	}

	result := db.Save(&categoryModel)
	if result.Error != nil {
		return nil, result.Error
	}

	category.CreatedAt = categoryModel.CreatedAt
	category.UpdatedAt = categoryModel.UpdatedAt

	out := r.Mapper.ToEntity(categoryModel)
	return out, nil
}

func (r *categoryRepository) Delete(ID int) error {
	db := r.DB

	categoryModel, err := r.FindByID(ID)
	if err != nil {
		return err
	}

	// Delete the category
	result := db.Delete(&categoryModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
