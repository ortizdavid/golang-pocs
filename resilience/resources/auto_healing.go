package resources

import (
	"context"
	"sync"
	"time"

	"github.com/ortizdavid/go-enterprise-micro/internal/infra/cache"
	"github.com/ortizdavid/go-enterprise-micro/internal/infra/database"
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

	// 1. Cria um contexto com timeout para a verificação
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 2. Se já existe e está respondendo, não faz nada
	if ah.res.Cache != nil && ah.res.Cache.Ping(ctx) == nil {
		return
	}

	// 3. Tenta recriar/reconectar o Cache
	newCache, err := cache.NewRedisCache( /* parâmetros de config */ )
	if err != nil {
		// Log de erro (caso tenha logger no res)
		return
	}

	// 4. Substitui pelo novo cache recuperado
	ah.res.Cache = newCache
}

func (ah *AutoHealing) RecoverDatabase() {
	ah.res.mu.Lock()
	defer ah.res.mu.Unlock()

	// 1. Cria um contexto com timeout para a verificação
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 2. Se já existe e está respondendo, não faz nada
	if ah.res.Database != nil && ah.res.Database.Ping(ctx) == nil {
		return
	}

	// 3. Tenta recriar/reconectar o Database
	newDb, err := database.GetDatabase( /* parâmetros de config */, ah.res.Logger)
	if err != nil {
		// Log de erro
		return
	}

	// 4. Substitui pela nova instância do banco recuperada
	ah.res.Database = newDb
}
