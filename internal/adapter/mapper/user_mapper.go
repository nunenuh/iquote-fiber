package mapper

import (
	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (qm *UserMapper) ToEntity(model *model.User) *entity.User {
	author := &entity.User{
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

func (qm *UserMapper) ToEntityWithPassword(model *model.User) *entity.User {
	author := &entity.User{
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

func (qm *UserMapper) ToEntityList(models []model.User) []*entity.User {
	out := make([]*entity.User, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}

func (qm *UserMapper) ToModel(entity *entity.User) *model.User {
	return &model.User{
		ID:          entity.ID,
		FullName:    entity.FullName,
		Email:       entity.Email,
		Username:    entity.Username,
		Password:    entity.Password,
		Phone:       entity.Phone,
		Level:       entity.Level,
		IsSuperuser: entity.IsSuperuser,
		IsActive:    true,
	}

}
