package main

import (
    "encoding/json"
    "html/template"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Product struct {
    Name  string
    Price float64
}

// Utility function to convert struct to JSON in the template
func toJSON(data interface{}) template.JS {
    a, _ := json.Marshal(data)
    return template.JS(a)
}

func main() {
    router := gin.Default()

    // Load HTML templates
    router.SetFuncMap(template.FuncMap{
        "tojson": toJSON,
    })
    router.LoadHTMLGlob("templates/*")

    // Sample products
    products := []Product{
        {"Product 1", 10.0},
        {"Product 2", 20.0},
        {"Product 3", 30.0},
        {"Product 4", 40.0},
        {"Product 5", 50.0},
        {"Product 6", 60.0},
        {"Product 7", 70.0},
        {"Product 8", 80.0},
        {"Product 9", 90.0},
        {"Product 10", 100.0},
    }

    // Endpoint to display multiple charts of products
    router.GET("/charts", func(c *gin.Context) {
        // Render the HTML template with the products
        c.HTML(http.StatusOK, "products.html", gin.H{
            "products": products,
        })
    })

    router.Run(":8080")
}
