package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/auth/usecase"
)

func ProvideAuthHandler(repo domain.IAuthRepository, svc domain.IAuthService) *AuthHandler {
	return NewAuthHandler(repo, svc)
}

type AuthHandler struct {
	repo    domain.IAuthRepository
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(repo domain.IAuthRepository, svc domain.IAuthService) *AuthHandler {
	return &AuthHandler{
		repo:    repo,
		usecase: usecase.NewAuthUsecase(repo, svc),
	}
}

func (h *AuthHandler) Register(route fiber.Router) {
	route.Post("/login", h.Login)
	route.Get("/verify", Protected(), h.VerifyToken)
	route.Get("/refresh", Protected(), h.RefreshToken)

}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	fmt.Printf("Login")

	// Get request body.
	request := &LoginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	token, err := h.usecase.Login(request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  token,
	})
}

func (h *AuthHandler) VerifyToken(ctx *fiber.Ctx) error {
	tokenString, err := h.getTokenString(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	claims, err := h.usecase.VerifyToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail token",
			"message": err.Error(),
		})
	}
	username := claims.Username

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Welcome to the application %s!", username),
	})
}

func (h *AuthHandler) RefreshToken(ctx *fiber.Ctx) error {
	tokenString, err := h.getTokenString(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	newTokenString, err := h.usecase.RefreshToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  newTokenString,
	})

}

func (h *AuthHandler) getTokenString(ctx *fiber.Ctx) (string, error) {
	authorizationString := ctx.Get("Authorization")
	parts := strings.Split(authorizationString, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Only Bearer Authorization are accepted!")
	}

	tokenString := parts[1]

	return tokenString, nil
}
