package cronjobs

import (
	"context"
	"log/slog"

	"github.com/ortizdavid/golang-pocs/cronjob-manager/services"
)

type BusinessJob struct {
	service *services.BusinessService
	logger *slog.Logger
}

func NewBusinessJob(service *services.BusinessService, logger *slog.Logger) *BusinessJob {
	return &BusinessJob{
		service: service,
		logger: logger,
	}
}

func (j *BusinessJob) Name() string {
	return "BusinessJob"
}

func (j *BusinessJob) Schedule() string {
	return "15s"
}

func (j *BusinessJob) Execute(ctx context.Context) error {
	j.logger.Info("Executing bunines task", "status", "started")

	// business logic call
	err := j.service.SomeBusinesOperation(ctx)
	if err != nil {
		j.logger.Error("Error Processing business", "error", err)
		return err
	}

	j.logger.Info("Businesss Logic Processed successfuly", "job", j.Name())
	return nil
}
