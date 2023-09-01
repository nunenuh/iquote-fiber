package webutils

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	// Other necessary imports...
)

// ContextAdapter is an interface to abstract common operations across different frameworks.
type ContextAdapter interface {
	GetParam(key string) string
	GetQueryParam(key string) string
	GetBody() ([]byte, error)
	SetResponse(code int, data interface{})
}

type FiberAdapter struct {
	Ctx *fiber.Ctx
}

func (fa *FiberAdapter) GetParam(key string) string {
	return fa.Ctx.Params(key)
}

func (fa *FiberAdapter) GetQueryParam(key string) string {
	return fa.Ctx.Query(key)
}

func (fa *FiberAdapter) GetBody() ([]byte, error) {
	return fa.Ctx.Body(), nil
}

func (fa *FiberAdapter) SetResponse(code int, data interface{}) {
	fa.Ctx.Status(code).JSON(data)
}

type GinAdapter struct {
	Ctx *gin.Context
}

func (ga *GinAdapter) GetParam(key string) string {
	return ga.Ctx.Param(key)
}

func (ga *GinAdapter) GetQueryParam(key string) string {
	return ga.Ctx.DefaultQuery(key, "")
}

func (ga *GinAdapter) GetBody() ([]byte, error) {
	body, err := ga.Ctx.GetRawData()
	return body, err
}

func (ga *GinAdapter) SetResponse(code int, data interface{}) {
	ga.Ctx.JSON(code, data)
}

// For usecases and repositories, you'd typically not work with the context directly.
// However, if there's a need to carry certain request-scoped data (e.g. User ID, Authorization Token)
// down to the usecase and repository layer, consider designing a custom struct or use context.Context.
