package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// load all static files
func ConfigStaticFiles(app *fiber.App) {
	app.Static("/", "./public/static")
}

// load all .html templates
func GetTemplateEngine() *html.Engine {
	engine := html.New("./public/templates", ".html")
	return engine
}
