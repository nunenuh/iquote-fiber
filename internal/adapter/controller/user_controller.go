package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nunenuh/iquote-fiber/internal/app/usecase"
	"github.com/nunenuh/iquote-fiber/internal/domain/entity"
	"github.com/nunenuh/iquote-fiber/internal/domain/repository"
)

type UserController struct {
	userRepository repository.IUserRepository
}

// Creates a new handler.
func NewUserController(route fiber.Router, userRepository repository.IUserRepository) {
	// Create a handler based on our created service / use-case.
	handler := &UserController{
		userRepository: userRepository,
	}
	// We will restrict this route with our JWT middleware.
	// You can inject other middlewares if you see fit here.
	// cityRoute.Use(auth.JWTMiddleware(), auth.GetDataFromJWT)

	// Routing for general routes.
	route.Get("/list", handler.GetAll)

	// Routing for specific routes.
	route.Get("/:userID", handler.GetByID)
	route.Post("/create", handler.Create)
	route.Patch("/:userID", handler.Update)
	route.Delete("/:userID", handler.Delete)

	// cityRoute.Put("/:cityID", handler.checkIfCityExistsMiddleware, handler.updateCity)
	// cityRoute.Delete("/:cityID", handler.checkIfCityExistsMiddleware, handler.deleteCity)
}

func (ctrl *UserController) GetByID(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(ctrl.userRepository)
	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	u, err := userUsecase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})
}

func (ctrl *UserController) GetAll(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(ctrl.userRepository)

	// Get the limit and offset query parameters, with default values
	limitStr := ctx.Query("limit", "10")  // Default limit to 10 if not provided
	offsetStr := ctx.Query("offset", "0") // Default offset to 0 if not provided

	// Convert them to integers
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid limit value"})
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid offset value"})
	}

	u, err := userUsecase.GetAll(limit, offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error fetching users"})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   u,
	})

}

func (ctrl *UserController) Create(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(ctrl.userRepository)
	// Define an instance of the User entity to hold the parsed request body
	var user entity.User

	// Parse the request body into the user instance
	if err := ctx.BodyParser(&user); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	// Use the usecase to create the user and handle any potential errors
	createdUser, err := userUsecase.Create(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	// Return the created user as a response
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   createdUser,
	})
}

func (ctrl *UserController) Update(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(ctrl.userRepository)

	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	user := entity.User{}

	if err := ctx.BodyParser(&user); err != nil {
		log.Printf("Parsing error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to parse request",
		})
	}

	// Use the usecase to create the user and handle any potential errors
	updatedUser, err := userUsecase.Update(id, &user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	// Return the created user as a response
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   updatedUser,
	})
}

func (ctrl *UserController) Delete(ctx *fiber.Ctx) error {
	userUsecase := usecase.NewUserUsecase(ctrl.userRepository)

	// Extract the user ID from the URL path parameter
	idStr := ctx.Params("userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID format",
		})
	}

	// Call the usecase's delete function
	err = userUsecase.Delete(id)
	if err != nil {
		log.Printf("Deletion error: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	// Return success message
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}
