package api

import (
	"fmt"
	"github.com/saas0503/factory-api/guard"
	"net/http"
)

type Controller struct {
	Name        string
	Middlewares []middleware
	mux         Mux
}

func NewController(name string) *Controller {
	return &Controller{
		Name: "/" + name,
	}
}

func (c *Controller) Use(middleware middleware) *Controller {
	c.Middlewares = append(c.Middlewares, middleware)
	return c
}

func (c *Controller) Auth() *Controller {
	c.Middlewares = append(c.Middlewares, guard.Authentication)
	return c
}

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
	route := fmt.Sprintf("%s %s%s", method, c.Name, path)

	mergeHandler := handler
	for _, m := range c.Middlewares {
		mergeHandler = m(mergeHandler)
	}

	c.mux[route] = mergeHandler
}
