package mapper

import (
	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type QuoteMapper interface {
	ToEntity(model *model.Quote) *entity.Quote
	ToModel(entity *entity.Quote) *model.Quote
	ToEntityList(models []model.Quote) []*entity.Quote
}

type quoteMapper struct{}

func NewQuoteMapper() QuoteMapper {
	return &quoteMapper{}
}

func (qm *quoteMapper) ToEntity(model *model.Quote) *entity.Quote {

	var categories []entity.Category
	for _, cat := range model.Categories {
		category := entity.Category{
			ID:   cat.ID,
			Name: cat.Name,
		}
		categories = append(categories, category)
	}

	var usersWhoLiked []entity.User
	for _, user := range model.UserWhoLiked {
		u := entity.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
		usersWhoLiked = append(usersWhoLiked, u)
	}

	authorEntity := entity.Author{
		ID:   model.Author.ID,
		Name: model.Author.Name,
	}

	quoteEntity := &entity.Quote{
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

func (qm *quoteMapper) ToModel(entity *entity.Quote) *model.Quote {

	var categories []model.Category
	for _, cat := range entity.Category {
		category := model.Category{
			ID:   cat.ID,
			Name: cat.Name,
		}
		categories = append(categories, category)
	}

	var usersWhoLiked []model.User
	for _, user := range entity.UserWhoLiked {
		u := model.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			// Add other necessary fields
		}
		usersWhoLiked = append(usersWhoLiked, u)
	}

	authorModel := model.Author{
		ID:   entity.Author.ID,
		Name: entity.Author.Name,
	}

	quoteModel := &model.Quote{
		ID:           entity.ID, // Assuming your quote model has an ID field
		QText:        entity.QText,
		Tags:         entity.Tags,
		AuthorID:     &entity.Author.ID,
		Author:       authorModel,
		Categories:   categories,
		UserWhoLiked: usersWhoLiked,
	}

	return quoteModel
}

func (qm *quoteMapper) ToEntityList(models []model.Quote) []*entity.Quote {
	out := make([]*entity.Quote, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}

// func (qm *quoteMapper) ToModel(entity *entity.Quote) *model.Quote {
// 	id, _ := strconv.Atoi(entity.ID)
// 	return &model.Quote{
// 		ID:    id,
// 		QText: entity.QText,
// 		Tags:  entity.Tags,
// 	}
// }

// func QuoteModelToEntity(quoteModel *model.Quote) *entity.Quote {
// 	out := &entity.Quote{
// 		ID:        strconv.Itoa(quoteModel.ID),
// 		QText:     quoteModel.QText,
// 		Tags:      quoteModel.Tags,
// 		CreatedAt: quoteModel.CreatedAt,
// 		UpdatedAt: quoteModel.UpdatedAt,
// 	}

// 	return out
// }

// func QuoteEntityToModel(quoteEntity *entity.Quote) *model.Quote {
// 	out := &model.Quote{
// 		ID:        strconv.Itoa(quoteEntity.ID),
// 		QText:     quoteEntity.QText,
// 		Tags:      quoteEntity.Tags,
// 		CreatedAt: quoteEntity.CreatedAt,
// 		UpdatedAt: quoteEntity.UpdatedAt,
// 	}

// 	return out
// }
