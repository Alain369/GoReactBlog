package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/go-fiber-blog/database"     

func main() {

	database.ConnectDB()

	// Start the server
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"message": "Welcome to my GoReact Blog!"})

		return c.SendString("Hello, World 👋!")

	})
	app.Listen(":8003")
}
