package business

import (
	"fmt"
	"log/slog"
)

type BusinessService struct {
	logger *slog.Logger
}

func NewBusinessService(logger *slog.Logger) *BusinessService {
	return &BusinessService{logger: logger}
}

func (s *BusinessService) ProcessOperation() error {
	fmt.Println("[BUSINES] - Processing specific business operation")
	// logic here

	s.logger.Info("Operation message processed")
	return  nil
}