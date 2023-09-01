package router

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/nunenuh/iquote-fiber/internal/auth/api"
	author "github.com/nunenuh/iquote-fiber/internal/author/api/v1"
	category "github.com/nunenuh/iquote-fiber/internal/category/api/v1"
	quote "github.com/nunenuh/iquote-fiber/internal/quote/api/v1"
	user "github.com/nunenuh/iquote-fiber/internal/user/api/v1"
	"go.uber.org/fx"
)

type AppHandler struct {
	fx.In

	App             *fiber.App
	AuthHandler     *auth.AuthHandler
	UserHandler     *user.UserHandler
	AuthorHandler   *author.AuthorHandler
	CategoryHandler *category.CategoryHandler
	QuoteHandler    *quote.QuoteHandler
}

func RegisterRoutes(m AppHandler) {
	m.AuthHandler.Register(m.App.Group("/api/v1/auth"))
	m.UserHandler.Register(m.App.Group("/api/v1/user"))
	m.AuthorHandler.Register(m.App.Group("/api/v1/author"))
	m.CategoryHandler.Register(m.App.Group("/api/v1/category"))
	m.QuoteHandler.Register(m.App.Group("/api/v1/quote"))
}
