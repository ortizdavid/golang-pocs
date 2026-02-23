package resources

import (
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/config"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/infra"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InfraResources centralize infra dependenciesl
type InfraResources struct {
    DB     *gorm.DB
    Cache  *redis.Client 
    Logger *zap.Logger
    Config *config.AppConfig
}

// NewInfraResources init needed connections
func NewInfraResources() *InfraResources {

    appConfig := config.LoadAppConfig()

    logger, _ := zap.NewProduction()

    db := infra.NewDatabase(appConfig.DbURL) 

    redisClient := infra.NewRedisClient(appConfig.RedisURL)

    return &InfraResources{
        DB:     db,
        Cache:  redisClient,
        Logger: logger,
        Config: appConfig,
    }
}