package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func main() {
    // RabbitMQ server details
    rabbitmqURL := "http://localhost:15672"
    username := "guest"
    password := "guest"

    // API endpoint for listing queues
    endpoint := rabbitmqURL + "/api/queues"

    // Create HTTP client
    client := &http.Client{}

    // Create HTTP request
    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        fmt.Println("Error creating HTTP request:", err)
        return
    }

    // Set basic authentication headers
    req.SetBasicAuth(username, password)

    // Send HTTP request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending HTTP request:", err)
        return
    }
    defer resp.Body.Close()

    // Read response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    // Parse JSON response
    var queues []map[string]interface{}
    err = json.Unmarshal(body, &queues)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    // Print queue names
    fmt.Println("Queues:")
    for _, queue := range queues {
        name := queue["name"].(string)
        fmt.Println(name)
    }
}
