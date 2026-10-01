package resources

import (
	"context"
	"log/slog"
	"sync"

	"github.com/ortizdavid/golang-pocs/resilience/infra/broker"
	"github.com/ortizdavid/golang-pocs/resilience/infra/cache"
	"github.com/ortizdavid/golang-pocs/resilience/infra/database"
	"github.com/ortizdavid/golang-pocs/resilience/ports"
)

type InfraResources struct {
	Database      ports.Database
	Cache         ports.Cache
	MessageBroker ports.MessageBroker
	mu            *sync.RWMutex
}

func NewInfraResources(logger *slog.Logger) (*InfraResources, error) {
	// Database - Postgres
	var db ports.Database
	dbClient, err := database.NewPostgresClient("postgres://user:password@127.0.0.1:5433/resilience_db?sslmode=disable")
	if err != nil {
		logger.Warn("Database offline, switching to No-Op", "error", err)
		db = database.NewDatabaseNoOp()
	} else {
		db = dbClient
	}

	// Cache - Redis
	var rdb ports.Cache
	cacheClient, err := cache.NewRedisCache("127.0.0.1:6380")
	if err != nil {
		logger.Warn("Redis offline, switching to No-Op", "error", err)
		rdb = cache.NewCacheNoOp()
	} else {
		rdb = cacheClient
	}

	// Broker - RabbitMQ
	var rbq ports.MessageBroker
	brokerClient, err := broker.NewRabbitMQClient("amqp://guest:guest@127.0.0.1:5673/")
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
		mu:            &sync.RWMutex{},
	}, nil
}

func (res *InfraResources) Close() error {
	if res.Database != nil {
		return res.Database.Close()
	}
	return nil
}

func (res *InfraResources) Ping(ctx context.Context) error {
	if res.Database != nil {
		return res.Database.Close()
	}
	return nil
}
