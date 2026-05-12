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
	router.Get("/health", h.health)
}

func (h *HealthHandler) health(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(c.UserContext(), 2*time.Second)
    defer cancel()

    // Ping
    dbErr := h.res.Database.Ping(ctx)
    cacheErr := h.res.Cache.Ping(ctx)
    brokerErr := h.res.MessageBroker.Ping(ctx)

    // Check status
    isDegraded := dbErr != nil || cacheErr != nil || brokerErr != nil

    statusData := fiber.Map{
        "status": "ok",
        "services": fiber.Map{
            "database": h.checkService(dbErr),
            "cache":    h.checkService(cacheErr),
            "broker":   h.checkService(brokerErr),
        },
        "timestamp": time.Now().Format(time.RFC3339),
    }

    if isDegraded {
        statusData["status"] = "degraded"
        return c.Status(fiber.StatusServiceUnavailable).JSON(statusData)
    }

    return c.JSON(statusData)
}

func (h *HealthHandler) checkService(err error) string {
    if err != nil {
        return "No-Op (Offline)"
    }
    return "Real (Online)"
}