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

type QuoteHandler struct {
	quoteRepository repository.IQuoteRepository
}

func NewQuoteHandler(quoteRepository repository.IQuoteRepository) *QuoteHandler {
	return &QuoteHandler{
		quoteRepository: quoteRepository,
	}
}

func (h *QuoteHandler) Register(route fiber.Router) {
	route.Use(middleware.Protected())
	route.Get("/list", h.GetAll)
	route.Get("/:quoteID", h.GetByID)
	route.Post("/create", h.Create)
	route.Patch("/:quoteID", h.Update)
	route.Delete("/:quoteID", h.Delete)
}

func ProvideQuoteHandler(repo repository.IQuoteRepository) *QuoteHandler {
	return NewQuoteHandler(repo)
}

func (h *QuoteHandler) GetByID(ctx *fiber.Ctx) error {
	quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)
	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	u, err := quoteUsecase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})
}

func (h *QuoteHandler) GetAll(ctx *fiber.Ctx) error {
	quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

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

	u, err := quoteUsecase.GetAll(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})

}

func (h *QuoteHandler) Create(ctx *fiber.Ctx) error {
	quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

	var quote entity.Quote

	if err := ctx.BodyParser(&quote); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	createdQuote, err := quoteUsecase.Create(&quote)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create quote",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   createdQuote,
	})
}

func (h *QuoteHandler) Update(ctx *fiber.Ctx) error {
	quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	quote := entity.Quote{}

	if err := ctx.BodyParser(&quote); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	updatedQuote, err := quoteUsecase.Update(id, &quote)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create quote",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   updatedQuote,
	})
}

func (h *QuoteHandler) Delete(ctx *fiber.Ctx) error {
	quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid quote ID format",
		})
	}

	err = quoteUsecase.Delete(id)
	if err != nil {
		log.Printf("Deletion error: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete quote",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Quote deleted successfully",
	})
}
