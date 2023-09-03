package quote

import (
	"strconv"
	"strings"

	author "github.com/nunenuh/iquote-fiber/internal/author/domain"
	category "github.com/nunenuh/iquote-fiber/internal/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/quote/domain"
)

type CreateQuoteRequest struct {
	QText      string   `json:"qtext"`
	Tags       []string `json:"tags"`
	AuthorID   string   `json:"author_id"`
	Categories []string `json:"categories"`
}

func (req *CreateQuoteRequest) ToEntity() (*domain.Quote, error) {
	authorID, err := strconv.Atoi(req.AuthorID)
	if err != nil {
		return nil, err
	}
	quote := &domain.Quote{
		QText:  req.QText,
		Tags:   strings.Join(req.Tags, ","),
		Author: author.Author{ID: authorID},
	}

	// Convert category IDs to domain.Category slice
	categories := make([]category.Category, len(req.Categories))
	for i, idStr := range req.Categories {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		categories[i] = category.Category{ID: id}
	}
	quote.Category = categories

	return quote, nil
}
