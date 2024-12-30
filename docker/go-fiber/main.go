package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	app.Get("/", func (ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	app.Get("/about", func (ctx *fiber.Ctx) error  {
		return ctx.SendString("Golang Fiber Docker Image")
	})

	app.Listen(":9000")

}