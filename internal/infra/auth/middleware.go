package auth

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nunenuh/iquote-fiber/internal/core/auth/domain"
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

func GenerateToken(auth *domain.Auth) (string, error) {
	dClaims := domain.CustomClaims{
		UserID:      auth.ID,
		Username:    auth.Username,
		Email:       auth.Email,
		IsSuperuser: auth.IsSuperuser,
	}

	b, err := json.Marshal(dClaims)
	if err != nil {
		return "", err
	}

	var jClaims map[string]any
	err = json.Unmarshal(b, &jClaims)
	if err != nil {
		return "", err
	}

	jClaims["exp"] = jwtExpire

	claims := jwt.MapClaims(jClaims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtExpire)).Unix()
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return newToken.SignedString([]byte(jwtSecret))
	}

	return "", errors.New("invalid token")
}
