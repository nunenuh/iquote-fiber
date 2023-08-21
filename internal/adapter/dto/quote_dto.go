package dto

import (
	"strings"

	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type CreateQuoteRequest struct {
	QText      string   `json:"qtext"`
	Tags       []string `json:"tags"`
	AuthorID   string   `json:"author_id"`
	Categories []string `json:"categories"`
}

func (req *CreateQuoteRequest) ToEntity() *entity.Quote {
	quote := &entity.Quote{
		QText:  req.QText,
		Tags:   strings.Join(req.Tags, ","),
		Author: entity.Author{ID: req.AuthorID},
	}

	// Convert category IDs to entity.Category slice
	categories := make([]entity.Category, len(req.Categories))
	for i, id := range req.Categories {
		categories[i] = entity.Category{ID: id}
	}
	quote.Category = categories

	return quote
}
