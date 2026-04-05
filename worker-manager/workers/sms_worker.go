package workers

import (
	"context"
	"log/slog"
	"time"

	"github.com/ortizdavid/golang-pocs/worker-manager/notification"
)

type SmsWorker struct {
	sms *notification.SmsNotification
}

func NewSmsWorker(sms *notification.SmsNotification) *SmsWorker {
	return &SmsWorker{sms: sms}
}

func (w *SmsWorker) Start(ctx context.Context) error {
	slog.Info("worker started", "name", w.Name())
	
	ticker := time.NewTicker(3 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            slog.Info("Worker stopping...")
            return ctx.Err()
        case <-ticker.C:
            if err := w.process(ctx); err != nil {
                slog.Error("Failed to process sms", "error", err)
            }
        }
    }
}

func (w *SmsWorker) Stop(ctx context.Context) error {
	return nil
}

func (w *SmsWorker) Name() string {
	return "[SmsWorker]"
}

func (w *SmsWorker) process(ctx context.Context) error {
	w.sms.Send(ctx, "00244936166699", "hello, its an sms worker")
	return nil
}