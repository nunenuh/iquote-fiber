package repository

import (
	"fmt"
	"strconv"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideCategoryRepository(db *gorm.DB) repository.ICategoryRepository {
	return NewCategoryRepository(db)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (r *categoryRepository) GetAll(limit int, offset int) ([]*entity.Category, error) {
	db := r.DB
	var categoryModel []model.Category
	result := db.Preload("Parent").Offset(offset).Limit(limit).Find(&categoryModel)
	if result.Error != nil {
		panic(result.Error)
	}

	out := make([]*entity.Category, 0)
	for _, u := range categoryModel {
		cat := &entity.Category{
			ID:          u.ID,
			Name:        u.Name,
			Description: u.Description,
			ParentID:    "",
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}
		if u.ParentID != nil {
			cat.ParentID = strconv.Itoa(*u.ParentID)
		} else {
			cat.ParentID = ""
		}
		out = append(out, cat)
	}
	return out, nil
}

func (r *categoryRepository) GetByID(ID int) (*entity.Category, error) {
	db := r.DB
	var categoryModel model.Category
	result := db.First(&categoryModel, ID)
	if result.Error != nil {
		panic(result.Error)
	}

	out := &entity.Category{
		ID:          categoryModel.ID,
		Name:        categoryModel.Name,
		Description: categoryModel.Description,
	}
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

	out := make([]*entity.Category, 0)
	for _, u := range categoryModel {
		out = append(out, &entity.Category{
			ID:          u.ID,
			Name:        u.Name,
			Description: u.Description,
			ParentID:    strconv.Itoa(*u.ParentID),
		})
	}
	return out, nil
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

	out := make([]*entity.Category, 0)
	for _, u := range categoryModel {
		out = append(out, &entity.Category{
			ID:          u.ID,
			Name:        u.Name,
			Description: u.Description,
			ParentID:    strconv.Itoa(*u.ParentID),
		})
	}
	return out, nil
}

func (r *categoryRepository) Create(category *entity.Category) (*entity.Category, error) {
	db := r.DB
	categoryModel := &model.Category{
		Name:        category.Name,
		Description: category.Description,
	}

	if category.ParentID != "" {
		var parentCategory model.Category
		parentResult := db.Where("id = ?", category.ParentID).First(&parentCategory)
		if parentResult.Error != nil {
			panic(parentResult.Error)
		} else {
			categoryModel.ParentID = &parentCategory.ID
		}
	}

	result := db.Create(&categoryModel)
	if result.Error != nil {
		panic(result.Error)
	}

	category.CreatedAt = categoryModel.CreatedAt
	category.UpdatedAt = categoryModel.UpdatedAt

	return category, nil
}

func (r *categoryRepository) Update(ID int, category *entity.Category) (*entity.Category, error) {
	db := r.DB
	var categoryModel model.Category
	resultFind := db.Where("id = ?", ID).First(&categoryModel)
	if resultFind.Error != nil {
		panic(resultFind.Error)
	} else {
		categoryModel.Name = category.Name
		categoryModel.Description = category.Description
	}

	if category.ParentID != "" {
		var parentCategory model.Category
		parentResult := db.Where("id = ?", category.ParentID).First(&parentCategory)
		if parentResult.Error != nil {
			panic(parentResult.Error)
		} else {
			categoryModel.ParentID = &parentCategory.ID
		}
	}

	result := db.Save(&categoryModel)
	if result.Error != nil {
		panic(result.Error)
	}
	return category, nil
}

func (r *categoryRepository) Delete(ID int) error {
	db := r.DB

	var categoryModel model.Category

	// Check if the category with the given ID exists
	if err := db.First(&categoryModel, ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("category with ID %d not found", ID)
		}
		return err
	}

	// Delete the category
	result := db.Delete(&categoryModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
