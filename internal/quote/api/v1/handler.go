package quote

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	auth "github.com/nunenuh/iquote-fiber/internal/auth/infra"

	"github.com/nunenuh/iquote-fiber/internal/quote/domain"
	"github.com/nunenuh/iquote-fiber/internal/quote/usecase"
	"github.com/nunenuh/iquote-fiber/internal/shared/param"
	"github.com/nunenuh/iquote-fiber/pkg/webutils"
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
	route.Get("/quote/name/:quoteName", h.GetByCategoryName)
	route.Get("/quote/id/:quoteID", h.GetByCategoryID)

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
		return ErrorInvalidQuoteIDFormat(ctx, err)
	}

	quote, err := h.usecase.GetByID(id)
	if err != nil {
		return ErrorFetchingQuote(ctx, err)
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Successfully get quote",
		quote,
	)
	return ctx.JSON(response)
}

func (h *QuoteHandler) GetAll(ctx *fiber.Ctx) error {
	p := new(param.Param)
	if err := ctx.QueryParser(p); err != nil {
		return ErrorInvalidPaginationParameters(ctx, err)
	}

	quote, err := h.usecase.GetAll(p)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Error fetching quote", err.Error())
	}

	pagination := webutils.NewPagination(p, len(quote))
	response := webutils.NewSuccessResponseWithPagination(quote, pagination)
	return ctx.JSON(response)
}

func (h *QuoteHandler) GetByAuthorID(ctx *fiber.Ctx) error {
	p := new(param.Param)
	if err := ctx.QueryParser(p); err != nil {
		return ErrorInvalidPaginationParameters(ctx, err)
	}

	idStr := ctx.Params("authorID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ErrorInvalidQuoteIDFormat(ctx, err)
	}

	quote, err := h.usecase.GetByAuthorID(id, p)
	if err != nil {
		return ErrorFetchingQuote(ctx, err)
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Like quote successful",
		quote,
	)
	return ctx.JSON(response)
}

func (h *QuoteHandler) Like(ctx *fiber.Ctx) error {
	quoteIDStr := ctx.Params("quoteID")
	quoteID, err := strconv.Atoi(quoteIDStr)
	if err != nil {
		log.Print(err)
		return ErrorInvalidQuoteIDFormat(ctx, err)
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["UserID"].(float64))

	quote, err := h.usecase.Like(quoteID, userID)
	if err != nil {
		log.Print(quote)
		return ErrorFetchingQuote(ctx, err)
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Like quote successful",
		quote,
	)
	return ctx.JSON(response)
}

func (h *QuoteHandler) Unlike(ctx *fiber.Ctx) error {
	quoteIDStr := ctx.Params("quoteID")
	quoteID, err := strconv.Atoi(quoteIDStr)
	if err != nil {
		return ErrorInvalidQuoteIDFormat(ctx, err)
	}

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["UserID"].(float64))

	quote, err := h.usecase.Unlike(quoteID, userID)
	if err != nil {
		return ErrorFetchingQuote(ctx, err)
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Unlike quote successful",
		quote,
	)
	return ctx.JSON(response)
}

func (h *QuoteHandler) GetByAuthorName(ctx *fiber.Ctx) error {
	name := ctx.Params("authorName")
	p := new(param.Param)
	if err := ctx.QueryParser(p); err != nil {
		return ErrorInvalidPaginationParameters(ctx, err)
	}

	quote, err := h.usecase.GetByAuthorName(name, p)
	if err != nil {
		return ErrorFetchingQuote(ctx, err)
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Unlike quote successful",
		quote,
	)
	return ctx.JSON(response)
}

func (h *QuoteHandler) GetByCategoryID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ErrorInvalidQuoteIDFormat(ctx, err)
	}
	p := new(param.Param)
	if err := ctx.QueryParser(p); err != nil {
		return ErrorInvalidPaginationParameters(ctx, err)
	}

	quote, err := h.usecase.GetByCategoryID(id, p)
	if err != nil {
		return ErrorFetchingQuote(ctx, err)
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Get By Category ID successful",
		quote,
	)
	return ctx.JSON(response)

}

func (h *QuoteHandler) GetByCategoryName(ctx *fiber.Ctx) error {
	name := ctx.Params("quoteName")

	p := new(param.Param)
	if err := ctx.QueryParser(p); err != nil {
		return ErrorInvalidPaginationParameters(ctx, err)
	}

	quote, err := h.usecase.GetByCategoryName(name, p)
	if err != nil {
		return ErrorFetchingQuote(ctx, err)
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Get By Category Name Successful",
		quote,
	)
	return ctx.JSON(response)

}

func (h *QuoteHandler) Create(ctx *fiber.Ctx) error {

	var quoteReq CreateQuoteRequest
	if err := ctx.BodyParser(quoteReq); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to parse request", err.Error())
	}

	quoteEntity, err := quoteReq.ToEntity()
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Failed to create quote", err.Error())
	}

	createdQuote, err := h.usecase.Create(quoteEntity)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Failed to create quote", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Successfully created quote",
		createdQuote,
	)
	return ctx.JSON(response)
}

func (h *QuoteHandler) Update(ctx *fiber.Ctx) error {
	// Get the ID from the URL parameter
	idParam := ctx.Params("quoteID")
	quoteID, err := strconv.Atoi(idParam)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Invalid quoteID format", err.Error())
	}

	// Parse the request body
	var quoteReq CreateQuoteRequest
	if err := ctx.BodyParser(&quoteReq); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to parse request", err.Error())
	}

	quoteEntity, err := quoteReq.ToEntity()
	if err != nil {
		return webutils.NewErrorResponse(
			ctx, fiber.StatusBadRequest, "Failed to update quote", err.Error(),
		)
	}

	// Use the usecase to update the quote
	updatedQuote, err := h.usecase.Update(quoteID, quoteEntity)
	if err != nil {
		return webutils.NewErrorResponse(
			ctx, fiber.StatusBadRequest, "Failed to update quote", err.Error(),
		)
	}

	response := webutils.NewSuccessResponseWithMessage(
		fmt.Sprintf("Successfully updated quote with ID:%d", quoteID),
		updatedQuote,
	)
	return ctx.JSON(response)
}

func (h *QuoteHandler) Delete(ctx *fiber.Ctx) error {

	idStr := ctx.Params("quoteID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ErrorInvalidQuoteIDFormat(ctx, err)
	}

	err = h.usecase.Delete(id)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to delete quote", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		fmt.Sprintf("Successfully deleted quote with ID:%d", id),
		nil,
	)
	return ctx.JSON(response)
}

func ErrorInvalidQuoteIDFormat(ctx *fiber.Ctx, err error) error {
	return webutils.NewErrorResponse(
		ctx,
		fiber.StatusBadRequest,
		"Invalid quoteID format",
		err.Error(),
	)
}

func ErrorFetchingQuote(ctx *fiber.Ctx, err error) error {
	return webutils.NewErrorResponse(
		ctx,
		fiber.StatusInternalServerError,
		"Error fetching quote",
		err.Error(),
	)
}

func ErrorInvalidPaginationParameters(ctx *fiber.Ctx, err error) error {
	return webutils.NewErrorResponse(
		ctx,
		fiber.StatusBadRequest,
		"Invalid pagination parameters",
		err.Error(),
	)
}
