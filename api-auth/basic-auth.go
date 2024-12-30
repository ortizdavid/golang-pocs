package main

import (
    "encoding/base64"
    "fmt"
    "net/http"
    "strings"
)

// User represents a user with username and password
type User struct {
    Username string
    Password string
}

// Define a list of users
var users = []User{
    {"user1", "pass1"},
    {"user2", "pass2"},
    // Add more users as needed
}

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get the Authorization header
        authHeader := r.Header.Get("Authorization")
        
        // Check if the Authorization header is set
        if authHeader == "" {
            // No Authorization header provided, request authentication
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Unauthorized access\n")
            return
        }
        
        // Check if the Authorization header starts with "Basic"
        if !strings.HasPrefix(authHeader, "Basic ") {
            // Invalid Authorization header
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprint(w, "Invalid Authorization header\n")
            return
        }
        
        // Decode the base64-encoded credentials
        credentials, err := base64.StdEncoding.DecodeString(authHeader[len("Basic "):])
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "Error decoding credentials: %v\n", err)
            return
        }
        
        // Split the credentials into username and password
        parts := strings.SplitN(string(credentials), ":", 2)
        if len(parts) != 2 {
            // Malformed credentials
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprint(w, "Malformed credentials\n")
            return
        }

        // Check if the provided username and password are valid
        username, password := parts[0], parts[1]
        if !isValidUser(username, password) {
            // Incorrect username or password
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprint(w, "Unauthorized access\n")
            return
        }
        
        // Call the original handler with the authenticated user
        handler(w, r)
    }
}

// isValidUser checks if the provided username and password are valid
func isValidUser(username, password string) bool {
    for _, user := range users {
        if user.Username == username && user.Password == password {
            return true
        }
    }
    return false
}

func main() {
    // Define your handler function
    helloHandler := func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, authenticated user!")
    }
    
    // Wrap your handler function with BasicAuth middleware
    http.HandleFunc("/", BasicAuth(helloHandler))
    
    // Start the HTTP server
    http.ListenAndServe(":8080", nil)
}
