package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/streadway/amqp"
)

// Message struct represents the data structure of the message
type Message struct {
    Text string `json:"text"`
    Number int `json:"number"`
    Boolean bool `json:"boolean"`
}

func main() {
    mux := http.NewServeMux()
    // Set up HTTP server and route
    mux.HandleFunc("POST /publish", publishHandler)

    // Start HTTP server
    log.Println("Starting HTTP server on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal("Failed to start HTTP server:", err)
    }
}

func publishHandler(w http.ResponseWriter, r *http.Request) {
    // Parse JSON request body
    var message Message
    if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Publish message to RabbitMQ
    if err := publishToRabbitMQ(message); err != nil {
        log.Println("Failed to publish message to RabbitMQ:", err)
        http.Error(w, "Failed to publish message", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Message published to RabbitMQ"))
}

func publishToRabbitMQ(message Message) error {
    // Connect to RabbitMQ server
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        return err
    }
    defer conn.Close()

    // Create a channel
    ch, err := conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    // Declare a queue
    q, err := ch.QueueDeclare(
        "golang_queue", // name
        false,     // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return err
    }

    // Convert message to JSON
    body, err := json.Marshal(message)
    if err != nil {
        return err
    }

    // Publish message to queue
    if err := ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    ); err != nil {
        return err
    }

    return nil
}
