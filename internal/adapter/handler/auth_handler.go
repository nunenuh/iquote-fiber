package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nunenuh/iquote-fiber/internal/adapter/common/hash"
	"github.com/nunenuh/iquote-fiber/internal/adapter/middleware"
	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

const secret = "asecret"

type AuthHandler struct {
	userRepository repository.IUserRepository
}

func NewAuthHandler(authRoute fiber.Router, userRepository repository.IUserRepository) {
	handler := &AuthHandler{
		userRepository: userRepository,
	}

	authRoute.Post("/login", handler.signInUser)
	authRoute.Get("/private", middleware.Protected(), handler.privateRoute)
}

func (h *AuthHandler) signInUser(c *fiber.Ctx) error {
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Get request body.
	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	userUsecase := usecase.NewUserUsecase(h.userRepository)
	u, err := userUsecase.GetByUsername(request.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Forbidden!",
		})
	}

	hPass, err := hash.HashPassword(request.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if request.Username != u.Username || hash.CheckHashPassword(hPass, u.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Wrong username or password!",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  signedToken,
	})
}

func (h *AuthHandler) privateRoute(ctx *fiber.Ctx) error {
	localUser := ctx.Locals("user")

	log.Printf("Type of localUser: %T\n", localUser)

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":   "success",
		"message":  "Welcome to the private route!",
		"username": username,
	})
}
