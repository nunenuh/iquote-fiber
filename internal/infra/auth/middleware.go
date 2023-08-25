package auth

import (
	"errors"
	"log"
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtHandler fiber.Handler
var jwtSecret string
var jwtExpire int64

type JWTUtility struct {
	secretKey  string
	jwtHandler fiber.Handler
}

func NewJWTUtility(secretKey string) *JWTUtility {
	return &JWTUtility{
		secretKey: secretKey,
	}
}
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
	log.Println(jwtExpire)

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

func GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = jwtExpire

	return token.SignedString([]byte(jwtSecret))
}

// func ParseToken(tokenStr string) (string, error) {
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(jwtSecret), nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		userId, ok := claims["sub"].(string)
// 		if !ok {
// 			return "", errors.New("invalid token claims")
// 		}
// 		return userId, nil
// 	}

// 	return "", errors.New("invalid token")
// }

func RefreshToken(token *jwt.Token) (string, error) {

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["exp"] = jwtExpire
		return token.SignedString([]byte(jwtSecret))
	}

	return "", errors.New("invalid token")
}
