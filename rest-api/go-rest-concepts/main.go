package main

import (
	"log"
	"net/http"

	"github.com/ortizdavid/go-rest-concepts/config"
	"github.com/ortizdavid/go-rest-concepts/handlers"
)

func main() {
	mux := http.NewServeMux()

	dbConn, _ := config.NewDBConnectionFromEnv("DATABASE_URL")
	
	handlers.RegisterRoutes(mux, dbConn.DB)

	log.Printf("Listenning on: http://%s", config.ListenAddr())
	http.ListenAndServe(config.ListenAddr(), mux)
}

