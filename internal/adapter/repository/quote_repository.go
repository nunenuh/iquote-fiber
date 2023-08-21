package repository

import (
	"fmt"
	"strconv"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideQuoteRepository(db *gorm.DB) repository.IQuoteRepository {
	return NewQuoteRepository(db)
}

type quoteRepository struct {
	DB *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) *quoteRepository {
	return &quoteRepository{
		DB: db,
	}
}

func (r *quoteRepository) GetAll(limit int, offset int) ([]*entity.Quote, error) {
	db := r.DB
	var quoteModel []model.Quote
	result := db.Preload("Author").Preload("Categories").Preload("UserWhoLiked").
		Offset(offset).Limit(limit).
		Find(&quoteModel)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get quote: %w", result.Error)
	}

	out := make([]*entity.Quote, 0)
	for _, u := range quoteModel {
		var categories []entity.Category
		for _, cat := range u.Categories {
			category := entity.Category{
				ID:   strconv.Itoa(cat.ID),
				Name: cat.Name,
			}
			categories = append(categories, category)
		}

		// Mapping users who liked the quote
		var usersWhoLiked []entity.User
		for _, user := range u.UserWhoLiked {
			u := entity.User{
				ID:       strconv.Itoa(user.ID),
				Username: user.Username,
				Email:    user.Email,
			}
			usersWhoLiked = append(usersWhoLiked, u)
		}

		authorEntity := entity.Author{ID: strconv.Itoa(u.Author.ID), Name: u.Author.Name}

		quoteEntity := &entity.Quote{
			ID:           strconv.Itoa(u.ID),
			QText:        u.QText,
			Tags:         u.Tags,
			Author:       authorEntity,
			Category:     categories,
			UserWhoLiked: usersWhoLiked,
			LikedCount:   len(usersWhoLiked),
			CreatedAt:    u.CreatedAt,
			UpdatedAt:    u.UpdatedAt,
		}

		out = append(out, quoteEntity)
	}
	return out, nil
}

func (r *quoteRepository) GetByID(ID int) (*entity.Quote, error) {
	db := r.DB
	var quoteModel model.Quote
	result := db.First(&quoteModel, ID)
	if result.Error != nil {
		panic(result.Error)
	}

	out := &entity.Quote{
		ID:        strconv.Itoa(quoteModel.ID),
		QText:     quoteModel.QText,
		Tags:      quoteModel.Tags,
		CreatedAt: quoteModel.CreatedAt,
		UpdatedAt: quoteModel.UpdatedAt,
	}
	return out, nil
}

func (r *quoteRepository) Create(quote *entity.Quote) (*entity.Quote, error) {
	db := r.DB

	var author model.Author
	authorID, err := strconv.Atoi(quote.Author.ID)
	if err != nil {
		return nil, fmt.Errorf("Author ID is not number!")
	}

	if err := db.First(&author, authorID).Error; err != nil {
		return nil, fmt.Errorf("Author with ID %s not found", quote.Author.ID)
	}

	quoteModel := &model.Quote{
		QText:    quote.QText,
		Tags:     quote.Tags,
		AuthorID: &author.ID,
	}

	ids := make([]int, len(quote.Category))
	for i, cat := range quote.Category {
		id, _ := strconv.Atoi(cat.ID)
		ids[i] = id
	}

	var categories []model.Category
	if err := db.Where("id IN ?", ids).Find(&categories).Error; err != nil {
		return nil, err // handle this error accordingly
	}

	quoteModel.Categories = categories

	// Create the quote.
	if err := db.Create(&quoteModel).Error; err != nil {
		return nil, err
	}

	// Update the entity fields from the model.
	quote.ID = strconv.Itoa(quoteModel.ID)
	quote.CreatedAt = quoteModel.CreatedAt
	quote.UpdatedAt = quoteModel.UpdatedAt

	return quote, nil
}

// Helper function added to the entity.Quote to extract category IDs.

func (r *quoteRepository) Update(ID int, quote *entity.Quote) (*entity.Quote, error) {
	db := r.DB
	quoteModel := &model.Quote{
		QText: quote.QText,
		Tags:  quote.Tags,
	}

	result := db.Save(&quoteModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return quote, nil
}

func (r *quoteRepository) Delete(ID int) error {
	db := r.DB

	var quoteModel model.Quote

	// Check if the quote with the given ID exists
	if err := db.First(&quoteModel, ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("quote with ID %d not found", ID)
		}
		return err
	}

	// Delete the quote
	result := db.Delete(&quoteModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
