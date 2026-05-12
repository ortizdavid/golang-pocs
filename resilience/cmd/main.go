package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/resilience/resources"
	"github.com/ortizdavid/golang-pocs/resilience/handlers" 
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	infra, err := resources.NewInfraResources(logger)
	if err != nil {
		logger.Error("Critical failure during resource initialization", "error", err)
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
		AppName:               "Go Resilience POCs",
	})

	handler := handlers.NewHealthHandler(infra)
	handler.Routes(app)

	go func() {
		logger.Info("Starting server on :3000")
		if err := app.Listen(":3000"); err != nil {
			logger.Error("Fiber server error", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit 
	
	logger.Info("Shutdown signal received. Cleaning up...")

	if err := app.Shutdown(); err != nil {
		logger.Error("Fiber shutdown failed", "error", err)
	}

	if err := infra.Close(); err != nil {
		logger.Error("Error closing infra resources", "error", err)
	}

	logger.Info("App terminated successfully.")
}