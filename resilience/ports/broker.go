package ports

import "context"

type MessageBroker interface {
    Publish(ctx context.Context, topic string, body []byte) error
    Consume(ctx context.Context, topic string) (<-chan []byte, error)
    Ping(ctx context.Context) error
}