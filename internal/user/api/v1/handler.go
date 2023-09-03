package user

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	auth "github.com/nunenuh/iquote-fiber/internal/auth/infra"
	"github.com/nunenuh/iquote-fiber/internal/shared/param"
	"github.com/nunenuh/iquote-fiber/internal/user/domain"
	"github.com/nunenuh/iquote-fiber/internal/user/usecase"
	"github.com/nunenuh/iquote-fiber/pkg/webutils"
)

type UserHandler struct {
	repo    domain.IUserRepository
	usecase *usecase.UserUsecase
}

func NewUserHandler(repo domain.IUserRepository) *UserHandler {
	return &UserHandler{
		repo:    repo,
		usecase: usecase.NewUserUsecase(repo),
	}
}

func (h *UserHandler) Register(route fiber.Router) {
	route.Use(auth.Protected())
	route.Get("/list", h.GetAll)
	route.Get("/:userID", h.GetByID)
	route.Post("/create", h.Create)
	route.Patch("/:userID", h.Update)
	route.Delete("/:userID", h.Delete)
}

func ProvideUserHandler(repo domain.IUserRepository) *UserHandler {
	return NewUserHandler(repo)
}

func (h *UserHandler) GetByID(ctx *fiber.Ctx) error {

	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Invalid userID format", err.Error())
	}

	user, err := h.usecase.GetByID(id)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Error fetching users", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Successfully get user",
		user,
	)
	return ctx.JSON(response)
}

func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	param := new(param.Param)
	if err := ctx.QueryParser(param); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Invalid pagination parameters", err.Error())
	}

	user, err := h.usecase.GetAll(param)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Error fetching users", err.Error())
	}

	pagination := webutils.NewPagination(param, len(user))
	response := webutils.NewSuccessResponseWithPagination(user, pagination)
	return ctx.JSON(response)

}

func (h *UserHandler) Create(ctx *fiber.Ctx) error {
	user := new(domain.User)
	if err := ctx.BodyParser(user); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to parse request", err.Error())
	}

	createdUser, err := h.usecase.Create(user)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusInternalServerError,
			"Failed to create user", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		"Successfully created user",
		createdUser,
	)
	return ctx.JSON(response)
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {

	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Invalid user ID format", err.Error())
	}

	user := new(domain.User)
	if err := ctx.BodyParser(user); err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to parse request", err.Error())
	}

	updatedUser, err := h.usecase.Update(id, user)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to create user", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		fmt.Sprintf("Successfully updated user with ID:%d", id),
		updatedUser,
	)
	return ctx.JSON(response)
}

func (h *UserHandler) Delete(ctx *fiber.Ctx) error {

	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Invalid user ID format", err.Error())
	}

	err = h.usecase.Delete(id)
	if err != nil {
		return webutils.NewErrorResponse(ctx, fiber.StatusBadRequest,
			"Failed to delete user", err.Error())
	}

	response := webutils.NewSuccessResponseWithMessage(
		fmt.Sprintf("Successfully deleted user with ID:%d", id),
		nil,
	)
	return ctx.JSON(response)
}
