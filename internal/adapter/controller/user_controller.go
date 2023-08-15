package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type UserController struct {
	userRepository repository.IUserRepository
}

// Creates a new handler.
func NewUserController(route fiber.Router, userRepository repository.IUserRepository) {
	// Create a handler based on our created service / use-case.
	handler := &UserController{
		userRepository: userRepository,
	}
	// We will restrict this route with our JWT middleware.
	// You can inject other middlewares if you see fit here.
	// cityRoute.Use(auth.JWTMiddleware(), auth.GetDataFromJWT)

	// Routing for general routes.
	// route.Get("", handler.GetByID)

	// Routing for specific routes.
	route.Get("/:userID", handler.GetByID)
	// cityRoute.Put("/:cityID", handler.checkIfCityExistsMiddleware, handler.updateCity)
	// cityRoute.Delete("/:cityID", handler.checkIfCityExistsMiddleware, handler.deleteCity)
}

func (ctrl *UserController) GetByID(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(ctrl.userRepository)
	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	u, err := userUsecase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})
}
