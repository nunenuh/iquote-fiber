package repository

import (
	"github.com/nunenuh/iquote-fiber/internal/adapter/database"
	"github.com/nunenuh/iquote-fiber/internal/adapter/database/model"
	"github.com/nunenuh/iquote-fiber/internal/domain"
)

type User struct{}

func (r User) Get(ID int) domain.User {
	db := database.Connection()
	var user model.User
	result := db.First(&user, ID)
	if result.Error != nil {
		panic(result.Error)
	}
	return domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

}
