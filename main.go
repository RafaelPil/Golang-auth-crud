package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(AuthDTO)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "User should have a email and password",
			})
		}

		fmt.Println(user)

		return c.JSON(user)
	})

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
