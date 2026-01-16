package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/go-translation/database"
	"github.com/ortizdavid/golang-pocs/go-translation/handlers"
	"github.com/ortizdavid/golang-pocs/go-translation/i18n"
	"github.com/ortizdavid/golang-pocs/go-translation/middlewares"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	app := fiber.New()

	// translation
	i18n.LoadTranslations()
	app.Use(middlewares.NewI18nMiddleware().Handle)

	db := database.InitDB("go_translation.db") // db
	defer db.Close()

	handlers.RegisterRoutes(app, db) // routes

	app.Listen(":8080")
}