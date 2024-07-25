package api

import (
	"fmt"
	"github.com/saas0503/factory-api/guard"
	"github.com/saas0503/factory-api/pipe"
	"net/http"
)

type Router struct {
	Name              string
	GlobalMiddlewares []middleware
	middlewares       []middleware
	/*	group             *Group */
	app *App
}

func (a *App) Group(name string, middlewares ...middleware) *Router {
	return &Router{
		Name:              name,
		GlobalMiddlewares: middlewares,
		middlewares:       []middleware{},
	}
}

func (r *Router) Version(version uint) *Router {
	v := fmt.Sprintf("/v%d", version)
	r.Name = r.Name + v
	return r
}

func (r *Router) Use(middleware middleware) *Router {
	r.middlewares = append(r.middlewares, middleware)
	return r
}

// Auth Guard

func (r *Router) Auth() *Router {
	r.middlewares = append(r.middlewares, guard.Authentication)
	return r
}

// Pagination

func (r *Router) Pagination() *Router {
	r.middlewares = append(r.middlewares, pipe.Pagination)
	return r
}

// Common Method

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle("GET", path, http.HandlerFunc(handler))
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle("POST", path, http.HandlerFunc(handler))
}

func (r *Router) Patch(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle("PATCH", path, http.HandlerFunc(handler))
}

func (r *Router) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle("PUT", path, http.HandlerFunc(handler))
}

func (r *Router) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.handle("DELETE", path, http.HandlerFunc(handler))
}

func (r *Router) handle(method string, path string, handler http.Handler) {
	route := fmt.Sprintf("%s %s%s", method, r.Name, IfSlashPrefixString(path))

	mergeHandler := handler

	for _, mg := range r.GlobalMiddlewares {
		mergeHandler = mg(mergeHandler)
	}

	for _, m := range r.middlewares {
		mergeHandler = m(mergeHandler)
	}

	r.middlewares = []middleware{}
	r.app.mux[route] = mergeHandler
}
