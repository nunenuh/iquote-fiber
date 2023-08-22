package dto

import (
	"strconv"
	"strings"

	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type CreateQuoteRequest struct {
	QText      string   `json:"qtext"`
	Tags       []string `json:"tags"`
	AuthorID   string   `json:"author_id"`
	Categories []string `json:"categories"`
}

func (req *CreateQuoteRequest) ToEntity() (*entity.Quote, error) {
	authorID, err := strconv.Atoi(req.AuthorID)
	if err != nil {
		return nil, err
	}
	quote := &entity.Quote{
		QText:  req.QText,
		Tags:   strings.Join(req.Tags, ","),
		Author: entity.Author{ID: authorID},
	}

	// Convert category IDs to entity.Category slice
	categories := make([]entity.Category, len(req.Categories))
	for i, idStr := range req.Categories {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		categories[i] = entity.Category{ID: id}
	}
	quote.Category = categories

	return quote, nil
}
