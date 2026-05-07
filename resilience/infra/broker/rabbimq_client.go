package broker

import (
	"context"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	conn *amqp091.Connection
	ch   *amqp091.Channel
}

func NewRabbitMQClient(url string) (*RabbitMQClient, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	return &RabbitMQClient{conn: conn, ch: ch}, err
}

func (r *RabbitMQClient) Publish(ctx context.Context, topic string, body []byte) error {
	return r.ch.PublishWithContext(ctx, "", topic, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}

func (r *RabbitMQClient) Consume(ctx context.Context, topic string) (<-chan []byte, error) {
	msgs, err := r.ch.Consume(topic, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	out := make(chan []byte)
	go func() {
		for d := range msgs {
			out <- d.Body
		}
	}()
	return out, nil
}

func (r *RabbitMQClient) Ping(ctx context.Context) error {
	if r.conn.IsClosed() {
		return amqp091.ErrClosed
	}
	return nil
}