package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type golangMessage struct {
	Text string `json:"text"`
    Number int `json:"number"`
    Boolean bool `json:"boolean"`
}

type dotnetMessage struct {
	Name string `json:"name"`
    Phone int `json:"phone"`
    Email string `json:"email"`
    Status bool `json:"status"`
}

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func main() {
    ConsumeQueue()
} 


func ConsumeQueue()  {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Change this to your RabbitMQ server address
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "golang_queue", // queue name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")

    forever := make(chan bool)

    go func() {
        for msg := range msgs {
			var obj golangMessage
			err := json.Unmarshal(msg.Body, &obj)
			if err != nil {
				fmt.Println(err)
				continue
			}
            log.Printf("Received a message: %s", msg.Body)
			log.Printf("Deserialized: %+v", obj)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever
}