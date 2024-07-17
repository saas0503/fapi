package guard

import (
	"context"
	"net/http"

	"github.com/saas0503/factory-api/config"
)

type UserContext string

const USER_CONTEXT UserContext = "user"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessToken string

		cfg, err := config.Load(".")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if accessToken = r.Header.Get("Authorization"); accessToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		payload, err := VerifyToken(accessToken, cfg.AccessTokenPublicKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		context := context.WithValue(r.Context(), USER_CONTEXT, payload)
		r = r.WithContext(context)

		next.ServeHTTP(w, r)
	})
}
