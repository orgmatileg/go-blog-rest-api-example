package handler

import (
	"fmt"
	"net/http"
)

// Ping Pong
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}
