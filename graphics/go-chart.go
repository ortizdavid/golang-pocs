package main

import (
    "github.com/gin-gonic/gin"
    "github.com/wcharczuk/go-chart/v2"
    "net/http"
)

type Product1 struct {
    Name  string
    Price float64
}

func main() {
    router := gin.Default()

    // Endpoint to generate chart
    router.GET("/chart", func(c *gin.Context) {
        // Sample products
        products := []Product1{
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

        // Prepare data for the chart
        xValues := []string{}
        yValues := []float64{}

        for _, product := range products {
            xValues = append(xValues, product.Name)
            yValues = append(yValues, product.Price)
        }

        graph := chart.BarChart{
            Height:   512,
            Width:    1024,
            BarWidth: 60,
            Bars: []chart.Value{},
        }

        for i, name := range xValues {
            graph.Bars = append(graph.Bars, chart.Value{
                Value: yValues[i],
                Label: name,
            })
        }

        c.Writer.Header().Set("Content-Type", "image/png")
        err := graph.Render(chart.PNG, c.Writer)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate chart"})
        }
    })

    router.Run(":8080")
}
