package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-htmlpdf-example/examples"
)


func main() {

	app := fiber.New()
	app.Get("/table-report", examples.TableReportHandler)
	app.Get("/simple-pdf", examples.SimplePdfHandler)
	app.Listen(":3000")
}