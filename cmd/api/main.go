package main

import (
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
	"github.com/nunenuh/iquote-fiber/internal/adapter/config"
	"github.com/nunenuh/iquote-fiber/internal/adapter/database"
	"github.com/nunenuh/iquote-fiber/internal/adapter/handler"
	"github.com/nunenuh/iquote-fiber/internal/adapter/repository"
)

func main() {

	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Configuration error: $s", err)

	}
	// 	// Try to connect to our database as the initial part.
	db, err := database.Connection(conf)
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}

	// Creates a new Fiber instance.
	app := fiber.New(fiber.Config{
		AppName:      "IQuote-Fiber Clean Arch",
		ServerHeader: "Fiber",
	})

	// Use global middlewares.
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

	userRepository := repository.NewUserRepository(db)

	handler.NewAuthHandler(app.Group("/api/v1/auth"), userRepository)
	handler.NewUserHandler(app.Group("/api/v1/user"), userRepository)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 8080.
	log.Fatal(app.Listen(":8080"))

}
