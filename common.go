package api

import (
	"encoding/json"
	"github.com/saas0503/factory-api/exception"
	"github.com/saas0503/factory-api/pipe"
	"net/http"
)

type middleware func(http.Handler) http.Handler

type Mux map[string]http.Handler

type ResponseOptions struct {
	Data    interface{}
	Total   int
	Message string
}

func TransformBody[P any](w http.ResponseWriter, r *http.Request) *P {
	var payload P

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		exception.ThrowInvalidRequest(w, err)
		return nil
	}

	errors := pipe.ValidateStruct(&payload)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(errors)
		if err != nil {
			return nil
		}
		return nil
	}

	return &payload
}

func JSON(w http.ResponseWriter, res ResponseOptions) {
	response := map[string]interface{}{
		"status": "success",
		"data":   res.Data,
	}

	if res.Total != 0 {
		response["total"] = res.Total
	}

	if res.Message != "" {
		response["message"] = res.Message
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
