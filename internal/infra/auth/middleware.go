package auth

import (
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var jwtHandler fiber.Handler
var jwtSecret string
var jwtExpire int64

func Protected() fiber.Handler {
	if jwtHandler == nil {
		panic("AuthMiddleware has not been initialized!")
	}
	return jwtHandler
}

func InitAuthMiddleware(secret string, expire string) {
	jwtSecret = secret

	exInt, err := strconv.Atoi(expire)
	if err != nil {
		panic(err)
	}

	jwtExpire = time.Now().Add(time.Minute * time.Duration(exInt)).Unix()
	jwtHandler = jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   jwtware.SigningKey{Key: []byte(jwtSecret)},
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
