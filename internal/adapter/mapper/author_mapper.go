package mapper

import (
	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type AuthorMapper struct{}

func NewAuthorMapper() *AuthorMapper {
	return &AuthorMapper{}
}

func (qm *AuthorMapper) ToEntity(model *model.Author) *entity.Author {
	author := &entity.Author{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
	return author
}

func (qm *AuthorMapper) ToEntityList(models []model.Author) []*entity.Author {
	out := make([]*entity.Author, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}

func (qm *AuthorMapper) ToModel(entity *entity.Author) *model.Author {
	return &model.Author{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		IsActive:    true,
	}

}
