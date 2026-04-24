package cronjobs

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/ortizdavid/golang-pocs/cronjob-manager/services"
)

type CronJobManager struct {
	Jobs   []BackgroundJob
	logger *slog.Logger
}

func NewCronJobManager(logger *slog.Logger) *CronJobManager {

	businesSvc := services.NewBusinessService(logger)
	businessJob := NewBusinessJob(businesSvc, logger)

	tempFileJob := NewTempFileJob("temp", logger)

	appCronJobs := []BackgroundJob{
		businessJob,
		tempFileJob,
	}

	return &CronJobManager{
		Jobs:   appCronJobs,
		logger: logger,
	}
}

func (m *CronJobManager) Run(ctx context.Context) error {
	for _, job := range m.Jobs {
		go m.start(ctx, job)
	}

	<-ctx.Done()
	m.logger.Info("Cron Engine shutting down...")
	return nil
}

func (m *CronJobManager) start(ctx context.Context, job BackgroundJob) {
	duration, err := time.ParseDuration(job.Schedule())
	if err != nil {
		m.logger.Error("Invalid schedule", "job", job.Name(), "error", err)
		return
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	m.logger.Info(fmt.Sprintf("Job [%s] scheduled every %s", job.Name(), duration))

	for {
		select {
		case <-ticker.C:
			m.execute(ctx, job)
		case <-ctx.Done():
			m.logger.Info("Job stopped", "name", job.Name())
			return
		}
	}
}

func (m *CronJobManager) execute(ctx context.Context, job BackgroundJob) {
	defer func() {
		if r := recover(); r != nil {
			m.logger.Error("Panic recovered", "job", job.Name(), "recover", r)
		}
	}()

	if err := job.Execute(ctx); err != nil {
		m.logger.Error("Execution failed", "job", job.Name(), "error", err)
	}
}

func (m *CronJobManager) Info() {
	fmt.Println("========================================================")
	fmt.Println("\t\t\tBACKGROUND JOBS")
	fmt.Println("========================================================")
	fmt.Println("STATUS         : RUNNING [MANAGER]")
	fmt.Println("PROTECTION     : RECOVER & GRACEFUL SHUTDOWN")

	for i, cj := range m.Jobs {
		fmt.Printf("[%d] %s\n", i+1, cj.Name())
	}
	fmt.Printf("========================================================\n\n")
}
