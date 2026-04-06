package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ortizdavid/golang-pocs/worker-manager/workers"
)

func main() {
	logger := slog.Default()

	// context for OS interruption (CTRL+C)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	manager := workers.NewWorkerManager(logger)
	manager.Info()

	go func() {
		if err := manager.Run(ctx); err != nil {
			logger.Error("Worker Manager exited with error", "error", err)
		}
	}()

	// Wait for Shutdown
	<-ctx.Done()
	logger.Info("Shutdown signal received. Starting graceful shutdown for Worker...")

	time.Sleep(2 * time.Second)

	logger.Info("Worker terminated successfully.")
}

