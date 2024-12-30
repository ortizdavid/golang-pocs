package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Golang Gin Docker Image",
        })
    })

	app.Run("localhost:7000")
}