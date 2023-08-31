package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/infra/auth"
)

func (h *AuthorHandler) Register(route fiber.Router) {
	route.Use(auth.Protected())
	route.Get("/list", h.GetAll)
	route.Get("/:authorID", h.GetByID)
	route.Post("/create", h.Create)
	route.Patch("/:authorID", h.Update)
	route.Delete("/:authorID", h.Delete)
}
