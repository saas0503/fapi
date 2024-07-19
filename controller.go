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

// Validate and transform

// func (c *Controller) ValidateBody(payload interface{}) *Controller {
//	c.Middlewares = append(c.Middlewares, pipe.Body[payload])
//	return c
// }

// Pagination

func (c *Controller) Pagination() *Controller {
	c.Middlewares = append(c.Middlewares, pipe.Pagination)
	return c
}

// Common Method

func (c *Controller) Get(path string, handler http.Handler) {
	c.Register("GET", path, handler)
}

func (c *Controller) Post(path string, handler http.Handler) {
	c.Register("POST", path, handler)
}

func (c *Controller) Patch(path string, handler http.Handler) {
	c.Register("PATCH", path, handler)
}

func (c *Controller) Put(path string, handler http.Handler) {
	c.Register("PUT", path, handler)
}

func (c *Controller) Delete(path string, handler http.Handler) {
	c.Register("DELETE", path, handler)
}

func (c *Controller) Register(method string, path string, handler http.Handler) {
	route := fmt.Sprintf("%s %s%s", method, c.Name, IfSlashPrefixString(path))

	mergeHandler := handler
	for _, m := range c.Middlewares {
		mergeHandler = m(mergeHandler)
	}

	c.Middlewares = []middleware{}
	c.mux[route] = mergeHandler
}
