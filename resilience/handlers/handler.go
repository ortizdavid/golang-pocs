package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/resilience/resources"
)

type HealthHandler struct {
	res *resources.InfraResources
}

func NewHealthHandler(res *resources.InfraResources) *HealthHandler {
	return &HealthHandler{res: res}
}

func (h *HealthHandler) Routes(router *fiber.App) {
	router.Get("/health")
}

func (h *HealthHandler) HealthCheck(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := h.res.Database.Ping(ctx)
	if err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  "degraded",
			"message": "Database is using No-Op mode",
		})
	}

	return c.JSON(fiber.Map{"status": "ok", "message": "Real Database is online"})
	
	
}