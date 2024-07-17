package api

import (
	"net/http"

	"github.com/saas0503/factory-api/guard"
)

type Router struct {
	Path        string
	Middlewares []middleware
	Handler     http.Handler
}

func (r *Router) Auth() *Router {
	r.Middlewares = append(r.Middlewares, guard.Authentication)
	return r
}
