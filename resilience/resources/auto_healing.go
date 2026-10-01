package resources

import (
	"context"
	"log/slog"
	"time"

	"github.com/ortizdavid/golang-pocs/resilience/infra/broker"
	"github.com/ortizdavid/golang-pocs/resilience/infra/cache"
	"github.com/ortizdavid/golang-pocs/resilience/infra/database"
)

type AutoHealing struct {
	res *InfraResources
}

func NewAutoHealing(res *InfraResources) *AutoHealing {
	return &AutoHealing{res: res}
}

func (ah *AutoHealing) StartAutohealing(ctx context.Context) {
	slog.Info("Starting background auto-healing service", "interval", "10s")
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				slog.Info("Stopping auto-healing service...")
				return
			case <-ticker.C:
				go ah.RecoverDatabase()
				go ah.RecoverCache()
				go ah.RecoverBroker()
			}
		}
	}()
}

func (ah *AutoHealing) RecoverCache() {
	ah.res.mu.Lock()
	defer ah.res.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if ah.res.Cache != nil {
		if err := ah.res.Cache.Ping(ctx); err == nil {
			return
		}
	}

	slog.Warn("Cache is offline or No-Op. Attempting reconnection...")
	cacheClient, err := cache.NewRedisCache("127.0.0.1:6380")
	if err != nil {
		slog.Error("Failed to reconnect Cache, keeping No-Op", "error", err)
		ah.res.Cache = cache.NewCacheNoOp()
		return
	}

	slog.Info("Cache recovered and swapped successfully!")
	ah.res.Cache = cacheClient
}

func (ah *AutoHealing) RecoverBroker() {
    ah.res.mu.Lock()
    defer ah.res.mu.Unlock()

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    if ah.res.MessageBroker != nil {
        if err := ah.res.MessageBroker.Ping(ctx); err == nil {
            return
        }
    }

    slog.Warn("Message Broker is offline or No-Op. Attempting reconnection...")
    brokerClient, err := broker.NewRabbitMQClient("amqp://guest:guest@127.0.0.1:5673/")
    if err != nil {
        slog.Error("Failed to reconnect Message Broker, keeping No-Op", "error", err)
        ah.res.MessageBroker = broker.NewBrokerNoOp()
        return
    }

    slog.Info("Message Broker recovered and swapped successfully!")
    ah.res.MessageBroker = brokerClient
}

func (ah *AutoHealing) RecoverDatabase() {
	ah.res.mu.Lock()
	defer ah.res.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if ah.res.Database != nil {
		if err := ah.res.Database.Ping(ctx); err == nil {
			return
		}
	}

	slog.Warn("Database is offline or No-Op. Attempting reconnection...")
	dbClient, err := database.NewPostgresClient("postgres://user:password@127.0.0.1:5433/resilience_db?sslmode=disable")
	if err != nil {
		slog.Error("Failed to reconnect Database, keeping No-Op", "error", err)
		ah.res.Database = database.NewDatabaseNoOp()
		return
	}

	slog.Info("Database recovered and swapped successfully!")
	ah.res.Database = dbClient
}