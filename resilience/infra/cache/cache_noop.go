package cache

import (
	"context"
	"fmt"
	"time"
)

type CacheNoOp struct{}

func NewCacheNoOp() *CacheNoOp {
    return &CacheNoOp{}
}

func (n *CacheNoOp) Get(ctx context.Context, key string) (any, error) {
    return nil, nil
}

func (n *CacheNoOp) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
    return nil 
}

func (n *CacheNoOp) Ping(ctx context.Context) error {
    return fmt.Errorf("cache service unavailable")
}