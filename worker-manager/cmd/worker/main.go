package main

import (
	"context"
	"log/slog"

	"github.com/ortizdavid/golang-pocs/worker-manager/workers"
)

func main() {
	logger := slog.Default()
	ctx := context.Background()

	manager := workers.NewWorkerManager(logger)
	if err := manager.Run(ctx); err != nil {

	}
}