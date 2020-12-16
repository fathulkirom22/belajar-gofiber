package router

import (
	"belajar-fiber/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes ...
func SetupRoutes(app *fiber.App) {

	// Api group
	api := app.Group("/api/v1")

	// Auth
	api.Post("/login", services.Login)

	// User
	user := api.Group("/user")
	user.Get("/", services.GetUsers)
	user.Get("/:id", services.GetUser)
	user.Post("/", services.NewUser)
	user.Delete("/:id", services.DeleteUser)
}
