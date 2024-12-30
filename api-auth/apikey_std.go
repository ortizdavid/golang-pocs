package main

import (
	"fmt"
	"net/http"
)

const (
	validApiKey = "key12345"
)

func ApiKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != validApiKey {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized. Invalid API Key")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have accessed the protected resource.")
}

func main() {
	mux := http.NewServeMux()

	// Attach the ApiKeyAuthMiddleware to the ProtectedHandler
	mux.Handle("/protected", ApiKeyAuthMiddleware(http.HandlerFunc(ProtectedHandler)))

	// Start the HTTP server
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
