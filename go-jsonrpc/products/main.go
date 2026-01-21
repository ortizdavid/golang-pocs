package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/ortizdavid/golang-pocs/go-jsonrpc/products/database"
	"github.com/ortizdavid/golang-pocs/go-jsonrpc/products/repositories"
	"github.com/ortizdavid/golang-pocs/go-jsonrpc/products/services"
)


func main() {

	db := database.InitDB("go_jsonrpc.db")
	defer db.Close()

	repo:= repositories.NewProductRepository(db)
	service := services.NewProductService(repo)

	// service registry
	server := rpc.NewServer()
	server.Register(service)

	// HTTP Hanlder to accept JSON-RPC
	http.HandleFunc("POST /rpc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		conn := struct{
			io.ReadCloser
			io.Writer
		}{r.Body, w}

		server.ServeCodec(jsonrpc.NewServerCodec(conn))
	})

	// Run Server
	log.Println("Product server running at http://localhost:1234/rpc")
	http.ListenAndServe(":1234", nil)
}