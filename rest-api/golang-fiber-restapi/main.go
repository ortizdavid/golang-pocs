package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-restapi/config"
	"github.com/ortizdavid/golang-fiber-restapi/entities"
	"github.com/ortizdavid/golang-fiber-restapi/controllers"
)

func main() {

	app := fiber.New()

	entities.SetupMigrations()
	controllers.SetupRoutes(app)
	app.Listen(config.ListenAndServe())
}