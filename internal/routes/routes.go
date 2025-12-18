package routes

import (
	"user-api/internal/handler"
	"user-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	// Global middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	api := app.Group("/users")

	api.Post("/", userHandler.CreateUser)
	api.Get("/:id", userHandler.GetUser)
	api.Get("/", userHandler.ListUsers)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)
}
