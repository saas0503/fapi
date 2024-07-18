package api

import "net/http"

type middleware func(http.Handler) http.Handler

type Mux map[string]http.Handler
