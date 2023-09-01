package api

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	auth "github.com/nunenuh/iquote-fiber/internal/auth/api"
	"github.com/nunenuh/iquote-fiber/internal/author/domain"
	"github.com/nunenuh/iquote-fiber/internal/author/usecase"
	"github.com/nunenuh/iquote-fiber/pkg/webutils"
)

type AuthorHandler struct {
	repo    domain.IAuthorRepository
	usecase *usecase.AuthorUsecase
}

func NewAuthorHandler(repo domain.IAuthorRepository) *AuthorHandler {
	return &AuthorHandler{
		repo:    repo,
		usecase: usecase.NewAuthorUsecase(repo),
	}
}

func ProvideAuthorHandler(repo domain.IAuthorRepository) *AuthorHandler {
	return NewAuthorHandler(repo)
}

func (h *AuthorHandler) Register(route fiber.Router) {
	route.Use(auth.Protected())
	route.Get("/list", h.GetAll)
	route.Get("/:authorID", h.GetByID)
	route.Post("/create", h.Create)
	route.Patch("/:authorID", h.Update)
	route.Delete("/:authorID", h.Delete)
}

func (h *AuthorHandler) GetByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	u, err := h.usecase.GetByID(id)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching author",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})
}

func (h *AuthorHandler) GetAll(ctx *fiber.Ctx) error {

	params, err := webutils.ParsePaginationParams(ctx)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError, "Invalid pagination parameters")
	}

	authors, err := h.usecase.GetAll(params)
	log.Print(authors)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError, "Error fetching authors")
	}

	pagination := webutils.NewPagination(params, int64(len(authors)))
	response := webutils.NewSuccessResponse(authors, pagination)

	return ctx.JSON(response)

}

func (h *AuthorHandler) Create(ctx *fiber.Ctx) error {
	var author domain.Author

	if err := ctx.BodyParser(&author); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	createdUser, err := h.usecase.Create(&author)
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
	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid author ID format",
		})
	}

	author := domain.Author{}

	if err := ctx.BodyParser(&author); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	updatedUser, err := h.usecase.Update(id, &author)
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
	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid author ID format",
		})
	}

	err = h.usecase.Delete(id)
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
