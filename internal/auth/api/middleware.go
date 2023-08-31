package auth

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var jwtHandler fiber.Handler

func Protected() fiber.Handler {
	if jwtHandler == nil {
		panic("AuthMiddleware has not been initialized!")
	}
	return jwtHandler
}

func InitAuthMiddleware(secret string) {
	jwtHandler = jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		// TokenLookup:  "header:Authorization",
	})

}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT!",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT!",
	})
}
