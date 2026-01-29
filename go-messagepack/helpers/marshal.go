package helpers

import (
	"io"
	"net/http"

	"github.com/vmihailenco/msgpack/v5"
)

// Onde o MessagePack é "chamado" de fato
func Unmarshal(r io.Reader, v interface{}) error {
	return msgpack.NewDecoder(r).Decode(v)
}

func Marshal(w io.Writer, v interface{}) error {
	return msgpack.NewEncoder(w).Encode(v)
}

func SendError(w http.ResponseWriter, message string, status int) {
    w.Header().Set("Content-Type", "application/x-msgpack")
    w.WriteHeader(status)
    Marshal(w, map[string]string{"error": message})
}
