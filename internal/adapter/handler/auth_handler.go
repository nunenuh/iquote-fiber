package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nunenuh/iquote-fiber/internal/adapter/dto"
	"github.com/nunenuh/iquote-fiber/internal/adapter/middleware"
	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

// const secret = "asecret"

type AuthHandler struct {
	repo repository.IUserRepository
}

func NewAuthHandler(userRepository repository.IUserRepository) *AuthHandler {
	return &AuthHandler{
		repo: userRepository,
	}
}

func (h *AuthHandler) Register(route fiber.Router) {
	route.Post("/login", h.signInUser)
	route.Get("/verify", middleware.Protected(), h.VerifyToken)
	route.Get("/refresh", middleware.Protected(), h.RefreshToken)

}

func ProvideAuthHandler(repo repository.IUserRepository) *AuthHandler {
	return NewAuthHandler(repo)
}

func (h *AuthHandler) signInUser(c *fiber.Ctx) error {

	// Get request body.
	request := &dto.LoginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	authUsecase := usecase.NewAuthUsecase(h.repo)
	user, err := authUsecase.Login(request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail use case",
			"message": err,
		})
	}

	token, err := middleware.GenerateToken(user.ID)
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

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Welcome to the application %s!", username),
	})
}

func (h *AuthHandler) RefreshToken(ctx *fiber.Ctx) error {
	localUser := ctx.Locals("user")

	log.Printf("Type of localUser: %T\n", localUser)

	token := ctx.Locals("user").(*jwt.Token)
	tokenStr, err := middleware.RefreshToken(token)
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
