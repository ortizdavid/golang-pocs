package services

import (
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/repositories"
	"go.uber.org/zap"
)

type UserService struct {
	repository *repositories.UserRepository
	logger *zap.Logger
}

func NewUserService(repo *repositories.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		repository: repo,
		logger: logger.Named("user"),
	}
}

// Business Operations for User