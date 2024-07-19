package pipe

import (
	"context"
	"encoding/json"
	"github.com/saas0503/factory-api/exception"
	"net/http"
)

type Token string

const PayloadToken Token = "payload"

func Body[P any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload P

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			exception.ThrowInvalidRequest(w, err)
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
