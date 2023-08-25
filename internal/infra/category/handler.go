package category

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/core/category/domain"
	"github.com/nunenuh/iquote-fiber/internal/core/category/usecase"
	"github.com/nunenuh/iquote-fiber/internal/infra/auth"
)

type CategoryHandler struct {
	categoryRepository domain.ICategoryRepository
}

func NewCategoryHandler(categoryRepository domain.ICategoryRepository) *CategoryHandler {
	return &CategoryHandler{
		categoryRepository: categoryRepository,
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

func ProvideCategoryHandler(repo domain.ICategoryRepository) *CategoryHandler {
	return NewCategoryHandler(repo)
}

func (h *CategoryHandler) GetByID(ctx *fiber.Ctx) error {
	categoryUsecase := usecase.NewCategoryUsecase(h.categoryRepository)
	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	u, err := categoryUsecase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})
}

func (h *CategoryHandler) GetAll(ctx *fiber.Ctx) error {
	categoryUsecase := usecase.NewCategoryUsecase(h.categoryRepository)

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

	u, err := categoryUsecase.GetAll(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching categorys",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})

}

func (h *CategoryHandler) Create(ctx *fiber.Ctx) error {
	categoryUsecase := usecase.NewCategoryUsecase(h.categoryRepository)

	var category domain.Category

	if err := ctx.BodyParser(&category); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	createdUser, err := categoryUsecase.Create(&category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create category",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   createdUser,
	})
}

func (h *CategoryHandler) Update(ctx *fiber.Ctx) error {
	categoryUsecase := usecase.NewCategoryUsecase(h.categoryRepository)

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

	updatedUser, err := categoryUsecase.Update(id, &category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create category",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   updatedUser,
	})
}

func (h *CategoryHandler) Delete(ctx *fiber.Ctx) error {
	categoryUsecase := usecase.NewCategoryUsecase(h.categoryRepository)

	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid category ID format",
		})
	}

	err = categoryUsecase.Delete(id)
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
		"message": "User deleted successfully",
	})
}
