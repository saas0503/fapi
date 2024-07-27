package guard

import (
	"context"
	"errors"
	"net/http"

	"github.com/saas0503/factory-api/exception"

	"github.com/saas0503/factory-api/config"
)

type Token string

const UserToken Token = "user"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessToken string

		cfg, err := config.Load(".")
		if err != nil {
			exception.ThrowInternalServerError(w, err)
			return
		}

		if accessToken = r.Header.Get("Authorization"); accessToken == "" {
			exception.ThrowTokenRequired(w, errors.New("missing Authorization header"))
			return
		}

		payload, err := VerifyToken(accessToken, cfg.AccessTokenPublicKey)
		if err != nil {
			exception.ThrowAuthFailed(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), UserToken, payload)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
