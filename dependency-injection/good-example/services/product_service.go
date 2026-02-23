package services

import (
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/repositories"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type ProductService struct {
	repository *repositories.ProductRepository
	cache *redis.Client 
	logger *zap.Logger
}

func NewProductService(repo *repositories.ProductRepository, cache *redis.Client, logger *zap.Logger) *ProductService {
	return &ProductService{
		repository: repo,
		cache: cache,
		logger: logger.Named("product"),
	}
}

// Business Operations for Product