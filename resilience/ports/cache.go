package ports

import (
	"context"
	"time"
)

type Cache interface {
    Get(ctx context.Context, key string) (any, error)
    Set(ctx context.Context, key string, value any, ttl time.Duration) error
    Ping(ctx context.Context) error
}