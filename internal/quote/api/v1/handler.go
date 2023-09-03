package quote

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	auth "github.com/nunenuh/iquote-fiber/internal/auth/infra"
	"github.com/nunenuh/iquote-fiber/internal/quote/domain"
	"github.com/nunenuh/iquote-fiber/internal/quote/usecase"
)

type QuoteHandler struct {
	repo    domain.IQuoteRepository
	usecase *usecase.QuoteUseCase
}

func NewQuoteHandler(repo domain.IQuoteRepository) *QuoteHandler {
	usecase := usecase.NewQuoteUsecase(repo)
	return &QuoteHandler{
		repo:    repo,
		usecase: usecase,
	}
}

func (h *QuoteHandler) Register(route fiber.Router) {
	route.Use(auth.Protected())
	route.Get("/list", h.GetAll)

	route.Get("/author/name/:authorName", h.GetByAuthorName)
	route.Get("/author/id/:authorID", h.GetByAuthorID)
	route.Get("/category/name/:categoryName", h.GetByCategoryName)
	route.Get("/category/id/:categoryID", h.GetByCategoryID)

	route.Get("/like/:quoteID", h.Like)
	route.Get("/unlike/:quoteID", h.Unlike)

	route.Get("/:quoteID", h.GetByID)
	route.Post("/create", h.Create)
	route.Patch("/:quoteID", h.Update)
	route.Delete("/:quoteID", h.Delete)
}

func ProvideQuoteHandler(repo domain.IQuoteRepository) *QuoteHandler {
	return NewQuoteHandler(repo)
}

func (h *QuoteHandler) GetByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	quote, err := h.usecase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   quote,
	})
}

func getQueryParam(ctx *fiber.Ctx, param string, defaultValue int) int {
	paramStr := ctx.Query(param, strconv.Itoa(defaultValue))
	paramInt, err := strconv.Atoi(paramStr)
	if err != nil {
		return defaultValue
	}
	return paramInt
}

func (h *QuoteHandler) getLimitOffset(ctx *fiber.Ctx) (int, int, error) {
	limit := getQueryParam(ctx, "limit", 10)
	offset := getQueryParam(ctx, "offset", 0)

	return limit, offset, nil
}

func (h *QuoteHandler) GetAll(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.repo)

	limit, offset, err := h.getLimitOffset(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	quote, err := h.usecase.GetAll(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   quote,
	})

}

func (h *QuoteHandler) GetByAuthorID(ctx *fiber.Ctx) error {

	limit, offset, err := h.getLimitOffset(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authorID",
		})
	}

	quote, err := h.usecase.GetByAuthorID(id, limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Like quote successful",
		"data":    quote,
	})
}

func (h *QuoteHandler) Like(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.repo)

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

	quote, err := h.usecase.Like(quoteID, userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Like quote successful",
		"data":    quote,
	})
}

func (h *QuoteHandler) Unlike(ctx *fiber.Ctx) error {
	// quoteUsecase := usecase.NewQuoteUsecase(h.repo)

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

	quote, err := h.usecase.Unlike(quoteID, userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Like quote successful",
		"data":    quote,
	})
}

func (h *QuoteHandler) GetByAuthorName(ctx *fiber.Ctx) error {
	name := ctx.Params("authorName")
	limit, offset, err := h.getLimitOffset(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	quote, err := h.usecase.GetByAuthorName(name, limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   quote,
	})

}

func (h *QuoteHandler) GetByCategoryID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid authorID",
		})
	}

	limit, offset, err := h.getLimitOffset(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	quote, err := h.usecase.GetByCategoryID(id, limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   quote,
	})

}

func (h *QuoteHandler) GetByCategoryName(ctx *fiber.Ctx) error {
	name := ctx.Params("categoryName")

	limit, offset, err := h.getLimitOffset(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	quote, err := h.usecase.GetByCategoryName(name, limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Error fetching quotes",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   quote,
	})

}

func (h *QuoteHandler) Create(ctx *fiber.Ctx) error {

	var quoteReq CreateQuoteRequest

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
	var quoteReq CreateQuoteRequest
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
