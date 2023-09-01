package author

import (
	"github.com/nunenuh/iquote-fiber/internal/author/domain"
	"github.com/nunenuh/iquote-fiber/internal/database/model"
)

type AuthorMapper struct{}

func NewAuthorMapper() *AuthorMapper {
	return &AuthorMapper{}
}

func (qm *AuthorMapper) ToEntity(model *model.Author) *domain.Author {
	author := &domain.Author{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
	return author
}

func (qm *AuthorMapper) ToEntityList(models []model.Author) []*domain.Author {
	out := make([]*domain.Author, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}

func (qm *AuthorMapper) ToModel(entity *domain.Author) *model.Author {
	return &model.Author{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		IsActive:    true,
	}

}
