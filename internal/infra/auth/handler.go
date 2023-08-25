package auth

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nunenuh/iquote-fiber/internal/core/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/core/auth/usecase"
	// "github.com/nunenuh/iquote-fiber/internal/infra/auth"
)

// const secret = "asecret"

type AuthHandler struct {
	repo    domain.IAuthorRepository
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(userRepository domain.IAuthorRepository) *AuthHandler {
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

func ProvideAuthHandler(repo domain.IAuthorRepository) *AuthHandler {
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

	user, err := h.usecase.Login(request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail use case",
			"message": err.Error(),
		})
	}

	token, err := GenerateToken(user.ID)
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
	localUser := ctx.Locals("user")

	log.Printf("Type of localUser: %T\n", localUser)

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	// expired := claims["exp"].(float64)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Welcome to the application %s!", username),
	})
}

func (h *AuthHandler) RefreshToken(ctx *fiber.Ctx) error {
	localUser := ctx.Locals("user")

	log.Printf("Type of localUser: %T\n", localUser)

	token := ctx.Locals("user").(*jwt.Token)
	tokenStr, err := RefreshToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail token",
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  tokenStr,
	})

}
