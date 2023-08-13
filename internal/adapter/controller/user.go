package controller

import (
	fiber "github.com/gofiber/fiber/v2/"
	"github.com/nunenuh/iquote-fiber/internal/adapter/repository"
	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
)

var (
	userRepository = repository.User{}
)

type Controller struct{}

func (ctr Controller) user(ctx *fiber.Ctx) {
	user := usecase.User(userRepository)
	ctx.JSON(user)
}
