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

	for _, worker := range wm.Workers {
		go wm.startWorker(ctx, worker, errChan)
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

 func (wm *WorkerManager) startWorker(ctx context.Context, worker BackgroundWorker, errChan chan error) {
	defer func() {
		if r := recover(); r != nil {
			wm.logger.Error(fmt.Sprintf("RECOVERED panic in worker %s: %v", worker.Name(), r), "", nil)
			errChan <- fmt.Errorf("worker %s panicked", worker.Name())
		}
	}()

	wm.logger.Info(fmt.Sprintf("Starting worker: %s", worker.Name()))

	if err := worker.Start(ctx); err != nil {
		errChan <- fmt.Errorf("worker %s failed: %w", worker.Name(), err)
	}
 }

func (wm *WorkerManager) Info() {
	fmt.Println("========================================================")
	fmt.Println("\t\t\tBACKGROUND WORKERS")
	fmt.Println("========================================================")
	fmt.Println("STATUS         : RUNNING [MANAGER]")
	fmt.Println("PROTECTION     : RECOVER & GRACEFUL SHUTDOWN")

	for i, w := range wm.Workers {
		fmt.Printf("[%d] %s\n", i+1, w.Name())
	}
	fmt.Printf("========================================================\n\n")
}