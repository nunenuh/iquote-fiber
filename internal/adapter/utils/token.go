package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nunenuh/iquote-fiber/internal/adapter/config"
)

func ProvideJWTUtility(cfg config.Configuration) *JWTUtility {
	return NewJWTUtility(cfg.JWTSecret)
}

type JWTUtility struct {
	secretKey string
}

func NewJWTUtility(secretKey string) *JWTUtility {
	return &JWTUtility{
		secretKey: secretKey,
	}
}

// GenerateToken generates a new JWT token.
func (j *JWTUtility) GenerateToken(userID string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(duration).Unix()

	return token.SignedString([]byte(j.secretKey))
}

// ParseToken parses the JWT token and returns the user ID.
func (j *JWTUtility) ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["sub"].(string)
		if !ok {
			return "", errors.New("invalid token claims")
		}
		return userId, nil
	}

	return "", errors.New("invalid token")
}

// VerifyToken checks if the token is valid.
func (j *JWTUtility) VerifyToken(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	return true, nil
}

// RefreshToken renews the token's expiration.
func (j *JWTUtility) RefreshToken(tokenStr string, duration time.Duration) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["exp"] = time.Now().Add(duration).Unix()
		return token.SignedString([]byte(j.secretKey))
	}

	return "", errors.New("invalid token")
}
