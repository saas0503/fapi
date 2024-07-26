package api

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/saas0503/factory-api/guard"
	"github.com/saas0503/factory-api/pipe"
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
	ParentMiddlewares []middleware
	middlewares       []middleware
	mux               Mux
}

func NewController(middlewares ...middleware) *Controller {
	return &Controller{
		ParentMiddlewares: middlewares,
		middlewares:       []middleware{},
		mux:               make(Mux),
	}
}

func (c *Controller) Registry(structs ...interface{}) {
	for _, item := range structs {
		ct := reflect.ValueOf(item).Elem()
		prefix := ct.Type().Name()
		prefix = strings.ToLower(prefix)
		prefix = strings.ReplaceAll(prefix, "controller", "")
		for i := 0; i < ct.NumField(); i++ {
			val := ct.Field(i)
			handler := val.Interface().(Handler)
			field := ct.Type().Field(i)
			auth := field.Tag.Get("guard")
			if auth == "authentication" {
				c.middlewares = append(c.middlewares, guard.Authentication)
			}
			pagination := field.Tag.Get("pagination")
			if pagination == "true" {
				c.middlewares = append(c.middlewares, pipe.Pagination)
			}
			if field.Tag.Get(string(GET)) != "" {
				c.register("GET", prefix, field.Tag.Get(string(GET)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(POST)) != "" {
				c.register("POST", prefix, field.Tag.Get(string(POST)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(PUT)) != "" {
				c.register("PUT", prefix, field.Tag.Get(string(PUT)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(PATCH)) != "" {
				c.register("PATCH", prefix, field.Tag.Get(string(PATCH)), http.HandlerFunc(handler))
			} else if field.Tag.Get(string(DELETE)) != "" {
				c.register("DELETE", prefix, field.Tag.Get(string(DELETE)), http.HandlerFunc(handler))
			}
		}
	}
}

func (c *Controller) register(method string, prefix string, path string, handler http.Handler) {
	route := fmt.Sprintf("%s %s%s", method, IfSlashPrefixString(prefix), IfSlashPrefixString(path))

	mergeHandler := handler

	for _, m := range c.ParentMiddlewares {
		mergeHandler = m(mergeHandler)
	}

	for _, m := range c.middlewares {
		mergeHandler = m(mergeHandler)
	}

	c.middlewares = []middleware{}
	c.mux[route] = mergeHandler
}
