package resources

import (
	"context"
	"sync"
	"time"

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
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				go ah.RecoverDatabase()
				go ah.RecoverCache()
			}
		}
	}()
}

func (ah *AutoHealing) RecoverCache() {
	ah.res.mu.Lock()
	defer ah.res.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Se o cache já estiver saudável (não for NoOp ou responder ao Ping se houver), podemos manter.
	// Como na sua POC você usa NoOp como fallback, verificamos se ele já está funcional:
	if ah.res.Cache != nil {
		// Se você tiver um método Ping na interface de Cache, pode validar aqui.
		// Exemplo: if ah.res.Cache.Ping(ctx) == nil { return }
	}

	// Tenta reconectar ao Redis
	cacheClient, err := cache.NewRedisCache("localhost:6380")
	if err != nil {
		// Se falhar a reconexão, mantém o No-Op
		ah.res.Cache = cache.NewCacheNoOp()
		return
	}

	// Sucesso na recuperação, substitui o recurso
	ah.res.Cache = cacheClient
}

func (ah *AutoHealing) RecoverDatabase() {
	ah.res.mu.Lock()
	defer ah.res.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if ah.res.Database != nil {
		// Se a porta Database tiver um Ping ou Check de saúde:
		// if err := ah.res.Database.Ping(ctx); err == nil { return }
	}

	// Tenta reconectar ao Postgres
	dbClient, err := database.NewPostgresClient("postgres://user:password@localhost:5433/resilience_db?sslmode=disable")
	if err != nil {
		// Se falhar, mantém o No-Op
		ah.res.Database = database.NewDatabaseNoOp()
		return
	}

	// Sucesso na recuperação, substitui o recurso
	ah.res.Database = dbClient
}
