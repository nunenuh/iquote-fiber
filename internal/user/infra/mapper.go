package user

import (
	"github.com/nunenuh/iquote-fiber/internal/database/model"
	"github.com/nunenuh/iquote-fiber/internal/user/domain"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (qm *UserMapper) ToEntity(model *model.User) *domain.User {
	author := &domain.User{
		ID:          model.ID,
		FullName:    model.FullName,
		Email:       model.Email,
		IsActive:    model.IsActive,
		Username:    model.Username,
		Phone:       model.Phone,
		Level:       model.Level,
		IsSuperuser: model.IsSuperuser,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
	return author
}

func (qm *UserMapper) ToEntityWithPassword(model *model.User) *domain.User {
	author := &domain.User{
		ID:          model.ID,
		FullName:    model.FullName,
		Email:       model.Email,
		Password:    model.Password,
		IsActive:    model.IsActive,
		Username:    model.Username,
		Phone:       model.Phone,
		Level:       model.Level,
		IsSuperuser: model.IsSuperuser,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
	return author
}

func (qm *UserMapper) ToEntityList(models []model.User) []*domain.User {
	out := make([]*domain.User, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}

func (qm *UserMapper) ToModel(domain *domain.User) *model.User {
	return &model.User{
		ID:          domain.ID,
		FullName:    domain.FullName,
		Email:       domain.Email,
		Username:    domain.Username,
		Password:    domain.Password,
		Phone:       domain.Phone,
		Level:       domain.Level,
		IsSuperuser: domain.IsSuperuser,
		IsActive:    true,
	}

}
