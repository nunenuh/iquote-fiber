package auth

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nunenuh/iquote-fiber/internal/core/auth/domain"
	"github.com/nunenuh/iquote-fiber/internal/infra/config"
)

func ProvideAuthService(cfg config.Configuration) domain.IAuthService {
	return NewAuthService(cfg)
}

type authService struct {
	Conf   config.Configuration
	Mapper *UserMapper
}

func NewAuthService(cfg config.Configuration) *authService {
	return &authService{
		Conf:   cfg,
		Mapper: NewUserMapper(),
	}
}

func (us *authService) GenerateToken(auth domain.Auth) (string, error) {
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

func (us *authService) VerifyToken(tokenString string) (*domain.CustomClaims, error) {
	token, err := us.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var cClaims domain.CustomClaims
		data, err := json.Marshal(claims)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(data, &cClaims); err != nil {
			return nil, err
		}

		return &cClaims, nil
	}

	return nil, errors.New("invalid token")
}

func (us *authService) RefreshToken(tokenString string) (string, error) {
	token, err := us.ParseToken(tokenString)
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

func (us *authService) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
	if err != nil {
		return nil, err
	}

	return token, nil
}
