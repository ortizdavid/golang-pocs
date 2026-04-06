package workers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/ortizdavid/golang-pocs/worker-manager/business"
)

type BusinessWorker struct {
	service *business.BusinessService
}

func NewBusinessWorker(service *business.BusinessService) *BusinessWorker {
	return &BusinessWorker{service: service}
}

func (w *BusinessWorker) Start(ctx context.Context) error {
	slog.Info(fmt.Sprintf("worker %s started",  w.Name()))
	
	ticker := time.NewTicker(3 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            slog.Info("Worker stopping...")
            return ctx.Err()
        case <-ticker.C:
            if err := w.process(ctx); err != nil {
                slog.Error("Failed to process business", "error", err)
            }
        }
    }
}

func (w *BusinessWorker) Stop(ctx context.Context) error {
	return nil
}

func (w *BusinessWorker) Name() string {
	return "[BusinesWorker]"
}

func (w *BusinessWorker) process(ctx context.Context) error {
	w.service.ProcessOperation(ctx)
	return nil
}