package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type middleware func(http.Handler) http.Handler

type Mux map[string]http.Handler

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("Received request:  %s %s %s", r.Method, r.URL.Path, elapsed)
	})
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	response := map[string]interface{}{
		"status":  "NotFound",
		"message": "The router not found",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
	return
}
