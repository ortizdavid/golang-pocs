package resources

import (
    "log/slog"

    "github.com/ortizdavid/golang-pocs/resilience/ports"
    "github.com/ortizdavid/golang-pocs/resilience/infra/database"
    "github.com/ortizdavid/golang-pocs/resilience/infra/cache"
    "github.com/ortizdavid/golang-pocs/resilience/infra/broker"
)

type InfraResources struct {
    Database      ports.Database
    Cache         ports.Cache
    MessageBroker ports.MessageBroker
}

func NewInfraResources(logger *slog.Logger) (*InfraResources, error) {

    // Database - Postgres
    var db ports.Database
    dbClient, err := database.NewPostgresClient("postgres://user:password@localhost:5432/resilience_db?sslmode=disable") 
    if err != nil {
        logger.Warn("Database offline, switching to No-Op", "error", err)
        db = database.NewDatabaseNoOp()
    } else {
        db = dbClient
    }

    // Cache - Redis
    var rdb ports.Cache
    cacheClient, err := cache.NewRedisCache("localhost:6379")
    if err != nil {
        logger.Warn("Redis offline, switching to No-Op", "error", err)
        rdb = cache.NewCacheNoOp()
    } else {
        rdb = cacheClient
    }

    // Broker - RabbitMQ
    var rbq ports.MessageBroker
    brokerClient, err := broker.NewRabbitMQClient("amqp://guest:guest@localhost:5672/")
    if err != nil {
        logger.Warn("RabbitMQ offline, switching to No-Op", "error", err)
        rbq = broker.NewBrokerNoOp()
    } else {
        rbq = brokerClient
    }

    return &InfraResources{
        Database:      db,
        Cache:         rdb,
        MessageBroker: rbq,
    }, nil
}

func (res *InfraResources) Close() error {
    if res.Database != nil {
        return res.Database.Close()
    }
    return nil
}