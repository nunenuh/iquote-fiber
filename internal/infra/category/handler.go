package category

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/core/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/core/category/usecase"
	"github.com/nunenuh/iquote-fiber/internal/infra/auth"
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid category ID",
		})
	}

	category, err := h.usecase.GetByID(id)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   category,
	})
}

func (h *CategoryHandler) GetAll(ctx *fiber.Ctx) error {
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

	categories, err := h.usecase.GetAll(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   categories,
	})

}

func (h *CategoryHandler) Create(ctx *fiber.Ctx) error {
	var category domain.Category

	if err := ctx.BodyParser(&category); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	createdCategory, err := h.usecase.Create(&category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create category",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   createdCategory,
	})
}

func (h *CategoryHandler) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	category := domain.Category{}

	if err := ctx.BodyParser(&category); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	updatedCategory, err := h.usecase.Update(id, &category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create category",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   updatedCategory,
	})
}

func (h *CategoryHandler) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid category ID format",
		})
	}

	err = h.usecase.Delete(id)
	if err != nil {
		log.Printf("Deletion error: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete category",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Category deleted successfully",
	})
}
