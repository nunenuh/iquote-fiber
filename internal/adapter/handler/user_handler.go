package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/adapter/middleware"
	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type UserHandler struct {
	userRepository repository.IUserRepository
}

func NewUserHandler(route fiber.Router, userRepository repository.IUserRepository) {

	handler := &UserHandler{
		userRepository: userRepository,
	}

	route.Use(middleware.Protected())
	route.Get("/list", handler.GetAll)
	route.Get("/:userID", handler.GetByID)
	route.Post("/create", handler.Create)
	route.Patch("/:userID", handler.Update)
	route.Delete("/:userID", handler.Delete)
}

func (h *UserHandler) GetByID(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(h.userRepository)
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

func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(h.userRepository)

	limitStr := ctx.Query("limit", "10")
	offsetStr := ctx.Query("offset", "0")

	// Convert them to integers
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid limit value",
		})
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid offset value",
		})
	}

	u, err := userUsecase.GetAll(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching users",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})

}

func (h *UserHandler) Create(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(h.userRepository)

	var user entity.User

	if err := ctx.BodyParser(&user); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	createdUser, err := userUsecase.Create(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   createdUser,
	})
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(h.userRepository)

	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	user := entity.User{}

	if err := ctx.BodyParser(&user); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	updatedUser, err := userUsecase.Update(id, &user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   updatedUser,
	})
}

func (h *UserHandler) Delete(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(h.userRepository)

	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID format",
		})
	}

	err = userUsecase.Delete(id)
	if err != nil {
		log.Printf("Deletion error: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}
