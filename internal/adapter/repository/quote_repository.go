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
		panic(result.Error)
	}

	out := make([]*entity.Quote, 0)
	for _, u := range quoteModel {
		cat := &entity.Quote{
			ID:        strconv.Itoa(u.ID),
			QText:     u.QText,
			Tags:      u.Tags,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}

		out = append(out, cat)
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
	quoteModel := &model.Quote{
		QText: quote.QText,
		Tags:  quote.Tags,
	}

	result := db.Create(&quoteModel)
	if result.Error != nil {
		panic(result.Error)
	}

	quote.CreatedAt = quoteModel.CreatedAt
	quote.UpdatedAt = quoteModel.UpdatedAt

	return quote, nil
}

func (r *quoteRepository) Update(ID int, quote *entity.Quote) (*entity.Quote, error) {
	db := r.DB
	quoteModel := &model.Quote{
		QText: quote.QText,
		Tags:  quote.Tags,
	}

	result := db.Save(&quoteModel)
	if result.Error != nil {
		panic(result.Error)
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
