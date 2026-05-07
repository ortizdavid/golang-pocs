package broker

import (
	"context"
	"fmt"
)

type BrokerNoOp struct{}

func NewBrokerNoOp() *BrokerNoOp {
    return &BrokerNoOp{}
}

func (n *BrokerNoOp) Publish(ctx context.Context, topic string, body []byte) error {
    return nil
}

func (n *BrokerNoOp) Consume(ctx context.Context, topic string) (<-chan []byte, error) {
    return make(chan []byte), nil 
}

func (n *BrokerNoOp) Ping(ctx context.Context) error {
    return fmt.Errorf("broker service unavailable (no-op mode)")
}