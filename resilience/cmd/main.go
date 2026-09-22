package main

import (
	"context" // <-- Adicionado para o contexto do AutoHealing
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/resilience/handlers"
	"github.com/ortizdavid/golang-pocs/resilience/resources"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	infra, err := resources.NewInfraResources(logger)
	if err != nil {
		logger.Error("Critical failure during resource initialization", "error", err)
		os.Exit(1)
	}

	// 1. Criamos um contexto raiz para gerenciar o ciclo de vida dos background workers (como o AutoHealing)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 2. Instancia e inicia o AutoHealing passando o contexto e os recursos de infra
	autoHealer := resources.NewAutoHealing(infra)
	autoHealer.StartAutohealing(ctx) // <-- Aqui ele começa a rodar em background

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

	// 3. Ao cancelar o contexto aqui, o loop do AutoHealing (select <-ctx.Done()) encerra instantaneamente
	cancel()

	if err := app.Shutdown(); err != nil {
		logger.Error("Fiber shutdown failed", "error", err)
	}

	if err := infra.Close(); err != nil {
		logger.Error("Error closing infra resources", "error", err)
	}

	logger.Info("App terminated successfully.")
}
