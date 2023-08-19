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

type AuthorHandler struct {
	authorRepository repository.IAuthorRepository
}

func NewAuthorHandler(authorRepository repository.IAuthorRepository) *AuthorHandler {
	return &AuthorHandler{
		authorRepository: authorRepository,
	}
}

func ProvideAuthorHandler(repo repository.IAuthorRepository) *AuthorHandler {
	return NewAuthorHandler(repo)
}

func (h *AuthorHandler) Register(route fiber.Router) {
	route.Use(middleware.Protected())
	route.Get("/list", h.GetAll)
	route.Get("/:authorID", h.GetByID)
	route.Post("/create", h.Create)
	route.Patch("/:authorID", h.Update)
	route.Delete("/:authorID", h.Delete)
}

func (h *AuthorHandler) GetByID(ctx *fiber.Ctx) error {
	authorUsecase := usecase.NewAuthorUsecase(h.authorRepository)
	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	u, err := authorUsecase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})
}

func (h *AuthorHandler) GetAll(ctx *fiber.Ctx) error {
	authorUsecase := usecase.NewAuthorUsecase(h.authorRepository)

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

	u, err := authorUsecase.GetAll(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching authors",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})

}

func (h *AuthorHandler) Create(ctx *fiber.Ctx) error {
	authorUsecase := usecase.NewAuthorUsecase(h.authorRepository)

	var author entity.Author

	if err := ctx.BodyParser(&author); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	createdUser, err := authorUsecase.Create(&author)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create author",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   createdUser,
	})
}

func (h *AuthorHandler) Update(ctx *fiber.Ctx) error {
	authorUsecase := usecase.NewAuthorUsecase(h.authorRepository)

	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	author := entity.Author{}

	if err := ctx.BodyParser(&author); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	updatedUser, err := authorUsecase.Update(id, &author)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create author",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   updatedUser,
	})
}

func (h *AuthorHandler) Delete(ctx *fiber.Ctx) error {
	authorUsecase := usecase.NewAuthorUsecase(h.authorRepository)

	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid author ID format",
		})
	}

	err = authorUsecase.Delete(id)
	if err != nil {
		log.Printf("Deletion error: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete author",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}
