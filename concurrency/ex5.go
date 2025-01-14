package main

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	OrderId      int `json:"order_id"`
	CustomerName string `json:"customer_name"`
	Price        float64 `json:"price"`
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8081", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	order := Order{OrderId: 1234, CustomerName: "John Smith", Price: 12.09}
	json.NewEncoder(w).Encode(order)
}