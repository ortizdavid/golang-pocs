package ports

import "context"

type Database interface {
    Ping(ctx context.Context) error
    Exec(ctx context.Context, query string, args ...any) error 
    Close() error
}