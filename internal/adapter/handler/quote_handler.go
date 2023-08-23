package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nunenuh/iquote-fiber/internal/adapter/dto"
	"github.com/nunenuh/iquote-fiber/internal/adapter/middleware"
	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type QuoteHandler struct {
	quoteRepository repository.IQuoteRepository
	usecase         *usecase.QuoteUseCase
}

func NewQuoteHandler(quoteRepository repository.IQuoteRepository) *QuoteHandler {
	usecase := usecase.NewQuoteUsecase(quoteRepository)
	return &QuoteHandler{
		quoteRepository: quoteRepository,
		usecase:         usecase,
	}
}

func (h *QuoteHandler) Register(route fiber.Router) {
	route.Use(middleware.Protected())
	route.Get("/list", h.GetAll)

	route.Get("/author/name/:authorName", h.GetByAuthorName)
	route.Get("/author/id/:authorID", h.GetByAuthorID)
	route.Get("/category/name/:categoryName", h.GetByCategoryName)
	route.Get("/category/id/:categoryID", h.GetByCategoryID)

	route.Get("/like/:quoteID", h.Like)

	route.Get("/:quoteID", h.GetByID)
	route.Post("/create", h.Create)
	route.Patch("/:quoteID", h.Update)
	route.Delete("/:quoteID", h.Delete)
}

func ProvideQuoteHandler(repo repository.IQuoteRepository) *QuoteHandler {
	return NewQuoteHandler(repo)
}

func (h *QuoteHandler) GetByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	u, err := h.usecase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})
}

func (h *QuoteHandler) GetAll(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

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

	u, err := h.usecase.GetAll(limit, offset)
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

func (h *QuoteHandler) GetByAuthorID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("authorID")
	limitStr := ctx.Query("limit", "10")
	offsetStr := ctx.Query("offset", "0")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authorID",
		})
	}

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

	u, err := h.usecase.GetByAuthorID(id, limit, offset)
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

func (h *QuoteHandler) Like(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

	quoteIDStr := ctx.Params("quoteID")
	quoteID, err := strconv.Atoi(quoteIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authorID",
		})
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	u, err := h.usecase.Like(quoteID, userID)
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

func (h *QuoteHandler) GetByAuthorName(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

	name := ctx.Params("authorName")
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

	u, err := h.usecase.GetByAuthorName(name, limit, offset)
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

func (h *QuoteHandler) GetByCategoryID(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

	idStr := ctx.Params("categoryID")
	limitStr := ctx.Query("limit", "10")
	offsetStr := ctx.Query("offset", "0")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authorID",
		})
	}

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

	u, err := h.usecase.GetByCategoryID(id, limit, offset)
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

func (h *QuoteHandler) GetByCategoryName(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.quoteRepository)

	name := ctx.Params("categoryName")
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

	u, err := h.usecase.GetByCategoryName(name, limit, offset)
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

	var quoteReq dto.CreateQuoteRequest

	if err := ctx.BodyParser(&quoteReq); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	quoteEntity, err := quoteReq.ToEntity()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create quote",
			"error":   err.Error(),
		})
	}

	createdQuote, err := h.usecase.Create(quoteEntity)
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
	// Get the ID from the URL parameter
	idParam := ctx.Params("quoteID")
	quoteID, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid quote ID format",
		})
	}

	// Parse the request body
	var quoteReq dto.CreateQuoteRequest
	if err := ctx.BodyParser(&quoteReq); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	quoteEntity, err := quoteReq.ToEntity()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create quote",
			"error":   err.Error(),
		})
	}

	// Use the usecase to update the quote
	updatedQuote, err := h.usecase.Update(quoteID, quoteEntity)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update quote",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   updatedQuote,
	})
}

func (h *QuoteHandler) Delete(ctx *fiber.Ctx) error {

	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid quote ID format",
		})
	}

	err = h.usecase.Delete(id)
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
