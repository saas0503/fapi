package router

import "net/http"

type Base struct {
	Path    string
	Handler http.Handler
}
