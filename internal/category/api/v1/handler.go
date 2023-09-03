package category

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	auth "github.com/nunenuh/iquote-fiber/internal/auth/infra"
	"github.com/nunenuh/iquote-fiber/internal/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/category/usecase"
	"github.com/nunenuh/iquote-fiber/internal/shared/param"
	"github.com/nunenuh/iquote-fiber/pkg/webutils"
)

func ProvideCategoryHandler(repo domain.ICategoryRepository) *CategoryHandler {
	return NewCategoryHandler(repo)
}

type CategoryHandler struct {
	repo    domain.ICategoryRepository
	usecase *usecase.CategoryUsecase
}

func NewCategoryHandler(repo domain.ICategoryRepository) *CategoryHandler {
	return &CategoryHandler{
		repo:    repo,
		usecase: usecase.NewCategoryUsecase(repo),
	}
}

func (h *CategoryHandler) Register(route fiber.Router) {
	route.Use(auth.Protected())
	route.Get("/list", h.GetAll)
	route.Get("/:categoryID", h.GetByID)
	// route.Get("/:parentID", handler.GetByParentID)
	route.Post("/create", h.Create)
	route.Patch("/:categoryID", h.Update)
	route.Delete("/:categoryID", h.Delete)
}

func (h *CategoryHandler) GetByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Invalid categoryID format", err.Error())
	}

	category, err := h.usecase.GetByID(id)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Error fetching category", err.Error())
	}
	response := webutils.NewSuccessResponseWithMessage(
		"Successfully get category",
		category,
	)
	return ctx.JSON(response)
}

func (h *CategoryHandler) GetAll(ctx *fiber.Ctx) error {
	p := new(param.Param)
	if err := ctx.QueryParser(p); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Invalid pagination parameters", err.Error())
	}

	category, err := h.usecase.GetAll(p)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Error fetching category", err.Error())
	}

	pagination := webutils.NewPagination(p, len(category))
	response := webutils.NewSuccessResponseWithPagination(category, pagination)

	return ctx.JSON(response)

}

func (h *CategoryHandler) Create(ctx *fiber.Ctx) error {
	category := new(domain.Category)
	if err := ctx.BodyParser(category); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to parse request", err.Error())
	}

	createdCategory, err := h.usecase.Create(category)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Failed to create category", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage("Successfully created category", createdCategory)
	return ctx.JSON(response)
}

func (h *CategoryHandler) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Invalid categoryID format", err.Error())
	}

	category := new(domain.Category)
	if err := ctx.BodyParser(&category); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to parse request", err.Error())
	}

	updatedCategory, err := h.usecase.Update(id, category)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to create category", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		fmt.Sprintf("Successfully updated category with ID:%d", id),
		updatedCategory,
	)
	return ctx.JSON(response)
}

func (h *CategoryHandler) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Invalid category ID format", err.Error())
	}

	err = h.usecase.Delete(id)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to delete category", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		fmt.Sprintf("Successfully deleted category with ID:%d", id),
		nil,
	)
	return ctx.JSON(response)
}
