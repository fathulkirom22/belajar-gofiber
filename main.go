package main

import (
	"belajar-fiber/config"
	"belajar-fiber/database"
	"belajar-fiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitConfig()
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)

	// 404 Handler
	app.Use(func(context *fiber.Ctx) error {
		return context.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
