package api

import (
	"encoding/json"
	"net/http"
)

type Req struct {
	*http.Request
}

type Res struct {
	http.ResponseWriter
}

type ResponseOptions struct {
	data    interface{}
	total   *int
	message *string
}

func JSON(w http.ResponseWriter, res ResponseOptions) {
	response := map[string]interface{}{
		"status":  "success",
		"data":    res.data,
		"message": *res.message,
		"total":   *res.total,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
