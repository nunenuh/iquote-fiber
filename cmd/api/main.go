package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nunenuh/iquote-fiber/internal/infra"
	"github.com/nunenuh/iquote-fiber/internal/infra/auth"
	"github.com/nunenuh/iquote-fiber/internal/infra/author"
	"github.com/nunenuh/iquote-fiber/internal/infra/category"
	"github.com/nunenuh/iquote-fiber/internal/infra/config"
	"github.com/nunenuh/iquote-fiber/internal/infra/database"
	"github.com/nunenuh/iquote-fiber/internal/infra/quote"
	"github.com/nunenuh/iquote-fiber/internal/infra/user"

	"go.uber.org/fx"
)

func createApp(config config.Configuration) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "IQuote Fiber Clean Arch",
		ServerHeader: "Fiber",
	})
	auth.InitAuthMiddleware(config.JWTSecret, config.JWTExpire)
	setupMiddleware(app)
	return app
}

func setupMiddleware(app *fiber.App) {
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
}

func startApp(lc fx.Lifecycle, app *fiber.App, config config.Configuration) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				// Prepare an endpoint for 'Not Found'.
				app.All("*", func(c *fiber.Ctx) error {
					errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
					return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
						"status":  "fail",
						"message": errorMessage,
					})
				})

				// Listen to port 8080.
				urls := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)
				log.Fatal(app.Listen(urls))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

func main() {

	fx.New(
		fx.Provide(config.ProvideConfig(".")),
		fx.Provide(database.ProvideDatabaseConnection()),
		fx.Provide(
			createApp,
			auth.ProvideAuthRepository,
			user.ProvideUserRepository,
			author.ProvideAuthorRepository,
			category.ProvideCategoryRepository,
			quote.ProvideQuoteRepository,

			auth.ProvideAuthHandler,
			user.ProvideUserHandler,
			author.ProvideAuthorHandler,
			category.ProvideCategoryHandler,
			quote.ProvideQuoteHandler,
		),
		fx.Invoke(
			infra.RegisterRoutes,
			startApp,
			// ... other invokable functions
		),
	).Run()

}
