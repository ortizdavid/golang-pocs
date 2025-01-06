package main

import (
	"golang-pocs/template-funcs/entities"
	"golang-pocs/template-funcs/helpers"

	"golang-pocs/template-funcs/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	engine := config.GetTemplateEngine() 
	helpers.AddTemplateFunc(engine)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	config.ConfigStaticFiles(app) 

	app.Get("", getCustomers)
	app.Get("/customers", getCustomers)
	app.Listen(":5000")
}

func getCustomers(c *fiber.Ctx) error {
	customerList := entities.CustomerList
	return c.Render("index", fiber.Map{
		"Title": "Customer List",
		"Customers": customerList,
		"Count": len(customerList),
	})
}