package repository

import (
	"fmt"

	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/adapter/mapper"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
	"gorm.io/gorm"
)

func ProvideQuoteRepository(db *gorm.DB) repository.IQuoteRepository {
	return NewQuoteRepository(db)
}

type quoteRepository struct {
	DB     *gorm.DB
	Mapper mapper.QuoteMapper
}

func NewQuoteRepository(db *gorm.DB) *quoteRepository {
	quoteMapper := mapper.NewQuoteMapper()
	return &quoteRepository{
		DB:     db,
		Mapper: quoteMapper,
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

	out := r.Mapper.ToEntityList(quoteModel)
	return out, nil
}

func (r *quoteRepository) FindByID(ID int) (*model.Quote, error) {
	db := r.DB
	var quoteModel model.Quote
	result := db.First(&quoteModel, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &quoteModel, nil
}

func (r *quoteRepository) FindAuthorByID(ID int) (*model.Author, error) {
	db := r.DB
	var author model.Author
	if err := db.First(&author, ID).Error; err != nil {
		return nil, fmt.Errorf("Author with ID %d not found", ID)
	}
	return &author, nil
}

func (r *quoteRepository) FindCategoriesByIDs(IDs []int) ([]model.Category, error) {
	db := r.DB
	var categories []model.Category
	if err := db.Where("id IN ?", IDs).Find(&categories).Error; err != nil {
		return nil, err // handle this error accordingly
	}
	return categories, nil
}

func (r *quoteRepository) GetByID(ID int) (*entity.Quote, error) {
	quoteModel, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	out := r.Mapper.ToEntity(quoteModel)
	return out, nil
}

func (r *quoteRepository) Create(quote *entity.Quote) (*entity.Quote, error) {
	db := r.DB

	author, err := r.FindAuthorByID(quote.Author.ID)
	if err != nil {
		return nil, err
	}

	quoteModel := &model.Quote{
		QText:    quote.QText,
		Tags:     quote.Tags,
		AuthorID: &author.ID,
		Author:   *author,
	}

	ids := make([]int, len(quote.Category))
	for i, cat := range quote.Category {
		ids[i] = cat.ID
	}

	categories, err := r.FindCategoriesByIDs(ids)
	if err != nil {
		return nil, err
	}

	quoteModel.Categories = categories

	if err := db.Create(&quoteModel).Error; err != nil {
		return nil, err
	}

	createdQuote := r.Mapper.ToEntity(quoteModel)
	return createdQuote, nil
}

// Helper function added to the entity.Quote to extract category IDs.
func (r *quoteRepository) Update(ID int, quote *entity.Quote) (*entity.Quote, error) {
	db := r.DB

	// Retrieve the existing quote
	existQuote, err := r.FindByID(ID)
	if err != nil {
		return nil, err
	}

	author, err := r.FindAuthorByID(quote.Author.ID)
	if err != nil {
		return nil, err
	}

	catIDs := make([]int, len(quote.Category))
	for i, cat := range quote.Category {
		catIDs[i] = cat.ID
	}

	categories, err := r.FindCategoriesByIDs(catIDs)
	if err != nil {
		return nil, err
	}

	existQuote.QText = quote.QText
	existQuote.Tags = quote.Tags
	existQuote.AuthorID = &author.ID
	existQuote.Author = *author
	existQuote.Categories = categories

	// Save the changes
	if err := db.Save(&existQuote).Error; err != nil {
		return nil, err
	}

	createdQuote := r.Mapper.ToEntity(existQuote)
	return createdQuote, nil
}

func (r *quoteRepository) Delete(ID int) error {
	db := r.DB
	quoteModel, err := r.FindByID(ID)
	if err != nil {
		return err
	}

	result := db.Delete(&quoteModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
