package infra

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/infra/auth"
	"github.com/nunenuh/iquote-fiber/internal/infra/author"
	"github.com/nunenuh/iquote-fiber/internal/infra/category"
	"github.com/nunenuh/iquote-fiber/internal/infra/quote"
	"github.com/nunenuh/iquote-fiber/internal/infra/user"
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
