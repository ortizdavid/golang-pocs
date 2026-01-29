package main

import (
	"bytes"
	"fmt"
	"net/http"
	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	baseURL := "http://localhost:8080/products"

	// 1. Testar CREATE
	newProd := map[string]interface{}{"code": "TEST-1", "name": "Produto Teste", "price": 99.9}
	body, _ := msgpack.Marshal(newProd)
	
	resp, _ := http.Post(baseURL, "application/x-msgpack", bytes.NewBuffer(body))
	fmt.Println("Create Status:", resp.Status)

	// 2. Testar GET ALL
	resp, _ = http.Get(baseURL)
	var products []map[string]interface{}
	msgpack.NewDecoder(resp.Body).Decode(&products)
	fmt.Printf("Total de produtos: %d\n", len(products))

	// 3. Testar UPDATE (com ID na rota)
	updateData := map[string]interface{}{"name": "Produto Atualizado"}
	updateBody, _ := msgpack.Marshal(updateData)
	
	req, _ := http.NewRequest(http.MethodPut, baseURL+"/1", bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/x-msgpack")
	resp, _ = http.DefaultClient.Do(req)
	fmt.Println("Update Status:", resp.Status)
}