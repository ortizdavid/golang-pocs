package database

import (
	"context"
	"fmt"
)

type DatabaseNoOp struct{}

func NewDatabaseNoOp() *DatabaseNoOp {
    return &DatabaseNoOp{}
}

func (n *DatabaseNoOp) Ping(ctx context.Context) error {
    return fmt.Errorf("database offline")
}

func (n *DatabaseNoOp) Exec(ctx context.Context, query string, args ...any) error {
    return fmt.Errorf("database in read-only/no-op mode") 
}

func (n *DatabaseNoOp) Close() error {
    return nil
}