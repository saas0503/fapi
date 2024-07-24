package interceptor

import (
	"encoding/json"
	"net/http"
)

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
