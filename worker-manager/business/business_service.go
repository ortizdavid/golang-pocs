package business

import (
	"context"
	"fmt"
	"log/slog"
)

type BusinessService struct {
	logger *slog.Logger
}

func NewBusinessService(logger *slog.Logger) *BusinessService {
	return &BusinessService{logger: logger}
}

func (s *BusinessService) ProcessOperation(ctx context.Context) error {
	fmt.Printf("\nProcessing specific business operation\n")
	// logic here

	s.logger.Info("Business operation processed")
	return  nil
}