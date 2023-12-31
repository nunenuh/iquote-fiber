package infra

import (
	"github.com/nunenuh/iquote-fiber/internal/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/database/model"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (qm *UserMapper) ToEntity(model *model.User) *domain.Auth {
	auth := &domain.Auth{
		ID:       model.ID,
		Email:    model.Email,
		Username: model.Username,
	}
	return auth
}

func (qm *UserMapper) ToEntityWithPassword(model *model.User) *domain.Auth {
	auth := &domain.Auth{
		ID:       model.ID,
		Email:    model.Email,
		Password: model.Password,
		Username: model.Username,
	}
	return auth
}

func (qm *UserMapper) ToEntityList(models []model.User) []*domain.Auth {
	out := make([]*domain.Auth, 0, len(models))

	for _, m := range models {
		out = append(out, qm.ToEntity(&m))
	}

	return out
}

func (qm *UserMapper) ToModel(domain *domain.Auth) *model.User {
	return &model.User{
		ID:       domain.ID,
		Email:    domain.Email,
		Username: domain.Username,
		Password: domain.Password,
		IsActive: true,
	}

}
