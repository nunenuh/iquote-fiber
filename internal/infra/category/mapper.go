package category

import (
	"github.com/nunenuh/iquote-fiber/internal/core/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/infra/database/model"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (qm *CategoryMapper) ToEntity(model *model.Category) *domain.Category {
	cat := &domain.Category{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	if model.Parent != nil {
		cat.ParentID = model.Parent.ID
		cat.Parent = &domain.Category{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
		}
	}
	return cat
}

func (qm *CategoryMapper) ToEntityList(models []model.Category) []*domain.Category {
	out := make([]*domain.Category, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}
