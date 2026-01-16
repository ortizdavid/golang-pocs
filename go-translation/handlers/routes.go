package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/go-translation/repositories"
)

func RegisterRoutes(router *fiber.App, db *sql.DB) {
	productRepo := repositories.NewProductRepository(db)

	NewProductHandler(productRepo).Routes(router)
}