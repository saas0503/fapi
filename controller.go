package api

import (
	"fmt"
	"github.com/saas0503/factory-api/guard"
	"github.com/saas0503/factory-api/pipe"
	"net/http"
	"reflect"
)

type Handler func(http.ResponseWriter, *http.Request)

type MethodKey string

const (
	GET    MethodKey = "GET"
	POST   MethodKey = "POST"
	PUT    MethodKey = "PUT"
	PATCH  MethodKey = "PATCH"
	DELETE MethodKey = "DELETE"
)

type Controller struct {
	Prefix      string
	Middlewares []middleware
	mux         Mux
}

func NewController(prefix string, middlewares ...middleware) *Controller {
	return &Controller{
		Prefix:      prefix,
		Middlewares: middlewares,
		mux:         make(Mux),
	}
}

func (c *Controller) Registry(structs ...interface{}) {
	for _, item := range structs {
		ct := reflect.ValueOf(item).Elem()
		for i := 0; i < ct.NumField(); i++ {
			val := ct.Field(i)
			handler := val.Interface().(Handler)
			field := ct.Type().Field(i)
			auth := field.Tag.Get("guard")
			if auth == "authentication" {
				c.Middlewares = append(c.Middlewares, guard.Authentication)
			}
			pagination := field.Tag.Get("pagination")
			if pagination == "true" {
				c.Middlewares = append(c.Middlewares, pipe.Pagination)
			}
			if field.Tag.Get(string(GET)) != "" {
				c.register("GET", field.Tag.Get(string(GET)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(POST)) != "" {
				c.register("POST", field.Tag.Get(string(POST)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(PUT)) != "" {
				c.register("PUT", field.Tag.Get(string(PUT)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(PATCH)) != "" {
				c.register("PATCH", field.Tag.Get(string(PATCH)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(DELETE)) != "" {
				c.register("DELETE", field.Tag.Get(string(DELETE)), http.HandlerFunc(handler))
			}
		}
	}
}

func (c *Controller) register(method string, path string, handler http.Handler) {
	route := fmt.Sprintf("%s %s%s", method, c.Prefix, IfSlashPrefixString(path))

	mergeHandler := handler

	for _, m := range c.Middlewares {
		mergeHandler = m(mergeHandler)
	}

	c.Middlewares = []middleware{}
	c.mux[route] = mergeHandler
}
