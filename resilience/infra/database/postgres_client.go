package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	db *sql.DB
}

func NewPostgresClient(dsn string) (*PostgresClient, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresClient{db: db}, nil
}

func (p *PostgresClient) Exec(ctx context.Context, query string, args ...any) error {
	_, err := p.db.ExecContext(ctx, query, args...)
	return err
}

func (p *PostgresClient) Ping(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func (p *PostgresClient) Close() error {
	return p.db.Close()
}