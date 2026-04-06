package workers

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/ortizdavid/golang-pocs/worker-manager/notification"
)

type EmailWorker struct {
	email *notification.EmailNotification
}

func NewEmailWorker(email *notification.EmailNotification) *EmailWorker {
	return &EmailWorker{email: email}
}

func (w *EmailWorker) Start(ctx context.Context) error {
	slog.Info(fmt.Sprintf("worker %s started",  w.Name()))
	
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            slog.Info("Worker stopping...")
            return ctx.Err()
        case <-ticker.C:
            if err := w.process(ctx); err != nil {
                slog.Error("Failed to process emails", "error", err)
            }
        }
    }
}

func (w *EmailWorker) Stop(ctx context.Context) error {
	return nil
}

func (w *EmailWorker) Name() string {
	return "[EmailWorker]"
}

func (w *EmailWorker) process(ctx context.Context) error {
	w.email.Send(ctx, "from@gmail.com", "to@gmail.com", "hello, its an email worker")
	return nil
}