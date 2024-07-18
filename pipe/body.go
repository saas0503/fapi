package pipe

import (
	"context"
	"encoding/json"
	"net/http"
)

type Token string

const PayloadToken Token = "payload"

func Body[P any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload P

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  "fail",
				"message": err.Error(),
			})
			if err != nil {
				return
			}
			return
		}

		errors := ValidateStruct(&payload)
		if errors != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(errors)
			if err != nil {
				return
			}
			return
		}

		ctx := context.WithValue(r.Context(), PayloadToken, payload)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
