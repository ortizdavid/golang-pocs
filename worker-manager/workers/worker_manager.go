package workers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ortizdavid/golang-pocs/worker-manager/business"
	"github.com/ortizdavid/golang-pocs/worker-manager/notification"
)

type WorkerManager struct {
	Workers []BackgroundWorker
	logger *slog.Logger
}

func NewWorkerManager(logger *slog.Logger) *WorkerManager {

	businessService := business.NewBusinessService(logger)
	businesWorker := NewBusinessWorker(businessService)

	emailNotif := notification.NewEmailNotification(logger)
	emailWorker := NewEmailWorker(emailNotif)

	smsNotif := notification.NewSmsNotification(logger)
	smsWorker := NewSmsWorker(smsNotif)

	appWorkers := []BackgroundWorker{
		businesWorker,
		emailWorker,
		smsWorker,
	}

	return &WorkerManager{
		Workers: appWorkers,
		logger: logger,
	}
}

func (wm *WorkerManager) Run(ctx context.Context) error {
	errChan := make(chan error, len(wm.Workers))

	for _, w := range wm.Workers {
		wm.logger.Info(fmt.Sprintf("Starting worker: %s", w.Name()))
		if err := w.Start(ctx); err != nil {
			errChan <- fmt.Errorf("worker %s failed: %w", w.Name(), err)
		}
	}

	select {
	case <-ctx.Done():
		wm.logger.Info("Shutdown signal received. Stopping workers...")
		return wm.StopAll(ctx)
	case err := <-errChan:
		wm.logger.Error("Critical worker failure", "error", err)
		return err
	}
}

func (wm *WorkerManager) StopAll(ctx context.Context) error {
	for _, w := range wm.Workers {
		if err := w.Stop(ctx); err != nil {
			wm.logger.Error(fmt.Sprintf("error stopping worker %s", w.Name()))
		}
	}
	return nil
}

func (wm *WorkerManager) Info() {
	fmt.Println("")
}