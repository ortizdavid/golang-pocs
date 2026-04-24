package services

import (
	"context"
	"log/slog"
	"time"
)

type BusinessService struct {
	logger *slog.Logger
}

func NewBusinessService(logger *slog.Logger) *BusinessService {
	return &BusinessService{
		logger: logger,
	}
}

func (s *BusinessService) SomeBusinesOperation(ctx context.Context) error {
	s.logger.Info("Executing a real busines operation ....")

	if err := s.saveToDatabase(ctx); err != nil {
        return err
    }
	
	s.logger.Info("Operation processed")
	return nil
}

func (s *BusinessService) saveToDatabase(ctx context.Context) error {
    s.logger.Info("Saving record to database...")
    
    time.Sleep(500 * time.Millisecond) 
    
    return nil
}