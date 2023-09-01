package quote

import (
	"github.com/nunenuh/iquote-fiber/internal/database/model"

	author "github.com/nunenuh/iquote-fiber/internal/author/domain"
	category "github.com/nunenuh/iquote-fiber/internal/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/quote/domain"
	user "github.com/nunenuh/iquote-fiber/internal/user/domain"
)

//	type IQuoteMapper interface {
//		ToEntity(model *model.Quote) *domain.Quote
//		ToModel(domain *domain.Quote) *model.Quote
//		ToEntityList(models []model.Quote) []*domain.Quote
//	}
type QuoteMapper struct{}

func NewQuoteMapper() *QuoteMapper {
	return &QuoteMapper{}
}

func (qm *QuoteMapper) ToEntity(model *model.Quote) *domain.Quote {

	var categories []category.Category
	for _, cat := range model.Categories {
		category := category.Category{
			ID:   cat.ID,
			Name: cat.Name,
		}
		categories = append(categories, category)
	}

	var usersWhoLiked []user.User
	for _, userdata := range model.UserWhoLiked {
		u := user.User{
			ID:       userdata.ID,
			Username: userdata.Username,
			Email:    userdata.Email,
		}
		usersWhoLiked = append(usersWhoLiked, u)
	}

	authorEntity := author.Author{
		ID:   model.Author.ID,
		Name: model.Author.Name,
	}

	quoteEntity := &domain.Quote{
		ID:           model.ID,
		QText:        model.QText,
		Tags:         model.Tags,
		Author:       authorEntity,
		Category:     categories,
		UserWhoLiked: usersWhoLiked,
		LikedCount:   len(usersWhoLiked),
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}

	return quoteEntity

}

func (qm *QuoteMapper) ToModel(domain *domain.Quote) *model.Quote {

	var categories []model.Category
	for _, cat := range domain.Category {
		category := model.Category{
			ID:   cat.ID,
			Name: cat.Name,
		}
		categories = append(categories, category)
	}

	var usersWhoLiked []model.User
	for _, user := range domain.UserWhoLiked {
		u := model.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			// Add other necessary fields
		}
		usersWhoLiked = append(usersWhoLiked, u)
	}

	authorModel := model.Author{
		ID:   domain.Author.ID,
		Name: domain.Author.Name,
	}

	quoteModel := &model.Quote{
		ID:           domain.ID, // Assuming your quote model has an ID field
		QText:        domain.QText,
		Tags:         domain.Tags,
		AuthorID:     &domain.Author.ID,
		Author:       authorModel,
		Categories:   categories,
		UserWhoLiked: usersWhoLiked,
	}

	return quoteModel
}

func (qm *QuoteMapper) ToEntityList(models []model.Quote) []*domain.Quote {
	out := make([]*domain.Quote, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}
