package main  

import (  
    "github.com/gofiber/fiber/v2"  
	"github.com/gofiber/fiber/v2/middleware/csrf"
)  

func main() {  
    app := fiber.New()  

    // Initialize CSRF middleware  
    app.Use(csrf.New(csrf.Config{  
		CookieSecure:  true, // Or use config.CsrfCookieSecure()
		CookieHTTPOnly: true,
		CookieName:     "csrf_token", 
        KeyLookup:    "header:X-CSRF-Token", // Lookup CSRF token from the header  
    }))  

    app.Get("/", func(c *fiber.Ctx) error {  
        token := c.Locals("csrf_token")
        return c.JSON(fiber.Map{  
            "csrf_token": token,  
        })  
    })  

    app.Post("/submit", func(c *fiber.Ctx) error {  
        return c.SendString("Form submitted!")  
    })  

    app.Listen(":9000")  
}