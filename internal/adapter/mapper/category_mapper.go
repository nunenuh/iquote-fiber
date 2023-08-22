package mapper

import (
	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (qm *CategoryMapper) ToEntity(model *model.Category) *entity.Category {
	cat := &entity.Category{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	if model.Parent != nil {
		cat.ParentID = model.Parent.ID
		cat.Parent = &entity.Category{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
		}
	}
	return cat
}

func (qm *CategoryMapper) ToEntityList(models []model.Category) []*entity.Category {
	out := make([]*entity.Category, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}
