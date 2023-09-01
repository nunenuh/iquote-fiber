package fiber

import "github.com/gofiber/fiber/v2"

type FiberContextAdapter struct {
	Ctx *fiber.Ctx
}

func (f FiberContextAdapter) QueryParam(name string) string {
	return f.Ctx.Query(name)
}

func (f FiberContextAdapter) JSON(statusCode int, v interface{}) error {
	return f.Ctx.Status(statusCode).JSON(v)
}
