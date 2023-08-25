package auth

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/core/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/core/auth/usecase"
)

type AuthHandler struct {
	repo    domain.IAuthRepository
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(userRepository domain.IAuthRepository) *AuthHandler {
	return &AuthHandler{
		repo:    userRepository,
		usecase: usecase.NewAuthUsecase(userRepository),
	}
}

func (h *AuthHandler) Register(route fiber.Router) {
	route.Post("/login", h.Login)
	route.Get("/verify", Protected(), h.VerifyToken)
	route.Get("/refresh", Protected(), h.RefreshToken)

}

func ProvideAuthHandler(repo domain.IAuthRepository) *AuthHandler {
	return NewAuthHandler(repo)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	// Get request body.
	request := &LoginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	auth, err := h.usecase.Login(request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail use case",
			"message": err.Error(),
		})
	}

	token, err := GenerateToken(auth)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail token",
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  token,
	})
}

func (h *AuthHandler) VerifyToken(ctx *fiber.Ctx) error {
	authorizationString := ctx.Get("Authorization")
	parts := strings.Split(authorizationString, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "error",
			"message": "Only Bearer Authorization are accepted!",
		})
	}
	tokenString := parts[1]
	// log.Printf("verify token:%s", tokenString)

	claims, err := VerifyToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail token",
			"message": err.Error(),
		})
	}
	username := claims["Username"].(string)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Welcome to the application %s!", username),
	})
}

func (h *AuthHandler) RefreshToken(ctx *fiber.Ctx) error {
	authorizationString := ctx.Get("Authorization")
	parts := strings.Split(authorizationString, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "error",
			"message": "Only Bearer Authorization are accepted!",
		})
	}
	tokenString := parts[1]
	// log.Printf("refresh token: %s", tokenString)
	newTokenString, err := RefreshToken(tokenString)
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
