package api

import (
	"fmt"
	"github.com/saas0503/factory-api/guard"
	"github.com/saas0503/factory-api/pipe"
	"net/http"
)

type Controller struct {
	Name        string
	Middlewares []middleware
	mux         Mux
}

func NewController(name string) *Controller {
	return &Controller{
		Name:        IfSlashPrefixString(name),
		Middlewares: []middleware{},
		mux:         make(Mux),
	}
}

func (c *Controller) Use(middleware middleware) *Controller {
	c.Middlewares = append(c.Middlewares, middleware)
	return c
}

// Auth Guard

func (c *Controller) Auth() *Controller {
	c.Middlewares = append(c.Middlewares, guard.Authentication)
	return c
}

// Pagination

func (c *Controller) Pagination() *Controller {
	c.Middlewares = append(c.Middlewares, pipe.Pagination)
	return c
}

// Common Method

func (c *Controller) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	c.Handling("GET", path, http.HandlerFunc(handler))
}

func (c *Controller) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	c.Handling("POST", path, http.HandlerFunc(handler))
}

func (c *Controller) Patch(path string, handler func(http.ResponseWriter, *http.Request)) {
	c.Handling("PATCH", path, http.HandlerFunc(handler))
}

func (c *Controller) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	c.Handling("PUT", path, http.HandlerFunc(handler))
}

func (c *Controller) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	c.Handling("DELETE", path, http.HandlerFunc(handler))
}

func (c *Controller) Handling(method string, path string, handler http.Handler) {
	route := fmt.Sprintf("%s %s%s", method, c.Name, IfSlashPrefixString(path))

	mergeHandler := handler
	for _, m := range c.Middlewares {
		mergeHandler = m(mergeHandler)
	}

	c.Middlewares = []middleware{}
	c.mux[route] = mergeHandler
}
