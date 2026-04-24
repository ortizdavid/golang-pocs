package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ortizdavid/golang-pocs/cronjob-manager/cronjobs"
)

func main() {
	logger := slog.Default()

	// context for OS interruption (CTRL+C)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	manager := cronjobs.NewCronJobManager(logger)
	manager.Info()

	go func() {
		if err := manager.Run(ctx); err != nil {
			logger.Error("CronJob Manager exited with error", "error", err)
		}
	}()

	// Wait for Shutdown
	<-ctx.Done()
	logger.Info("Shutdown signal received. Starting graceful shutdown for CronJob...")

	time.Sleep(2 * time.Second)

	logger.Info("Cron Job terminated successfully.")
}
