package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	auth "github.com/nunenuh/iquote-fiber/internal/auth/api"
	"github.com/nunenuh/iquote-fiber/internal/author/domain"
	"github.com/nunenuh/iquote-fiber/internal/author/usecase"
	"github.com/nunenuh/iquote-fiber/internal/shared/param"
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
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest, "Invalid authorID format", err.Error())
	}

	author, err := h.usecase.GetByID(id)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError, "Error fetching authors", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Successfully get author",
		author,
	)
	return ctx.JSON(response)
}

func (h *AuthorHandler) GetAll(ctx *fiber.Ctx) error {
	param := new(param.Param)
	if err := ctx.QueryParser(param); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError, "Invalid pagination parameters", err.Error())
	}

	authors, err := h.usecase.GetAll(param)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError, "Error fetching authors", err.Error())
	}

	pagination := webutils.NewPagination(param, len(authors))
	response := webutils.NewSuccessResponseWithPagination(authors, pagination)

	return ctx.JSON(response)
}

func (h *AuthorHandler) Create(ctx *fiber.Ctx) error {
	author := new(domain.Author)
	if err := ctx.BodyParser(author); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest, "Failed to parse request", err.Error())
	}

	createdAuthor, err := h.usecase.Create(author)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to create author", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage("Successfully created author", createdAuthor)
	return ctx.JSON(response)
}

func (h *AuthorHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("authorID"))
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest, "Invalid author ID format", err.Error())
	}

	var author domain.Author
	if err := ctx.BodyParser(&author); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest, "Failed to parse request", err.Error())
	}

	updatedAuthor, err := h.usecase.Update(id, &author)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest, "Failed to create author", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		fmt.Sprintf("Successfully updated author with ID:%d", id),
		updatedAuthor,
	)
	return ctx.JSON(response)
}

func (h *AuthorHandler) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest, "Invalid author ID format", err.Error())
	}

	err = h.usecase.Delete(id)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest, "Failed to delete author", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(fmt.Sprintf("Successfully deleted author with ID:%d", id), nil)
	return ctx.JSON(response)
}
