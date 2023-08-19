package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AppHandler struct {
	fx.In

	App             *fiber.App
	AuthHandler     *AuthHandler
	UserHandler     *UserHandler
	AuthorHandler   *AuthorHandler
	CategoryHandler *CategoryHandler
	QuoteHandler    *QuoteHandler
}

func RegisterRoutes(m AppHandler) {
	m.AuthHandler.Register(m.App.Group("/api/v1/auth"))
	m.UserHandler.Register(m.App.Group("/api/v1/user"))
	m.AuthorHandler.Register(m.App.Group("/api/v1/author"))
	m.CategoryHandler.Register(m.App.Group("/api/v1/category"))
	m.QuoteHandler.Register(m.App.Group("/api/v1/quote"))
}
