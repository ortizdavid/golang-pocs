package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const DEFAULT_API_KEY = "key123"

func ApiKeyAuthMiddleware2(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")

	if apiKey == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized. API Key missing")
	}

	if apiKey != DEFAULT_API_KEY {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized. Invalid API Key")
	}

	return c.Next()
}

func main() {
	app := fiber.New()

	// Middleware registration for specific routes
	app.Use("/protected", ApiKeyAuthMiddleware2)

	// Route definition
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Protected route
	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("This is a protected endpoint")
	})

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
