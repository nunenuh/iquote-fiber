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

	authApi "github.com/nunenuh/iquote-fiber/internal/auth/api/v1"
	authInfra "github.com/nunenuh/iquote-fiber/internal/auth/infra"
	"github.com/nunenuh/iquote-fiber/internal/router"

	authorApi "github.com/nunenuh/iquote-fiber/internal/author/api/v1"
	authorInfra "github.com/nunenuh/iquote-fiber/internal/author/infra"

	categoryApi "github.com/nunenuh/iquote-fiber/internal/category/api/v1"
	categoryInfra "github.com/nunenuh/iquote-fiber/internal/category/infra"

	"github.com/nunenuh/iquote-fiber/internal/database"
	quoteApi "github.com/nunenuh/iquote-fiber/internal/quote/api/v1"
	quoteInfra "github.com/nunenuh/iquote-fiber/internal/quote/infra"

	userApi "github.com/nunenuh/iquote-fiber/internal/user/api/v1"
	userInfra "github.com/nunenuh/iquote-fiber/internal/user/infra"

	"github.com/nunenuh/iquote-fiber/pkg/config"

	"go.uber.org/fx"
)

func createApp(config config.Configuration) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "IQuote Fiber Clean Arch",
		ServerHeader: "Fiber",
	})
	authInfra.InitAuthMiddleware(config.JWTSecret)
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
			authInfra.ProvideAuthRepository,
			userInfra.ProvideUserRepository,
			authorInfra.ProvideAuthorRepository,
			categoryInfra.ProvideCategoryRepository,
			quoteInfra.ProvideQuoteRepository,

			authInfra.ProvideAuthService,

			authApi.ProvideAuthHandler,
			userApi.ProvideUserHandler,
			authorApi.ProvideAuthorHandler,
			categoryApi.ProvideCategoryHandler,
			quoteApi.ProvideQuoteHandler,
		),
		fx.Invoke(
			router.RegisterRoutes,
			startApp,
			// ... other invokable functions
		),
	).Run()

}
