package api

import (
	"errors"
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

type BaseController struct {
	Prefix      string
	Middlewares []middleware
}

func Registry(structs ...interface{}) Mux {
	var mux = make(Mux)

	for _, item := range structs {
		var GlobalMiddlewares []middleware

		// Get Base
		base := reflect.ValueOf(item).Elem().FieldByName("Base").Interface().(BaseController)
		prefix := base.Prefix
		GlobalMiddlewares = base.Middlewares

		ct := reflect.ValueOf(item).Elem()
		for i := 0; i < ct.NumField(); i++ {
			field := ct.Type().Field(i)

			if field.Name == "Base" {
				continue
			}

			// Get middlewares
			var middlewares []middleware

			// Middlewares: check guard auth
			auth := field.Tag.Get("guard")
			if auth == "authentication" {
				middlewares = append(middlewares, guard.Authentication)
			}
			// Middlewares: check pagination
			pagination := field.Tag.Get("pagination")
			if pagination == "true" {
				middlewares = append(middlewares, pipe.Pagination)
			}

			// Get route path
			var routers []string
			if field.Tag.Get(string(GET)) != "" {
				routers = append(routers, "GET", field.Tag.Get(string(GET)))
			} else if field.Tag.Get(string(POST)) != "" {
				routers = append(routers, "POST", field.Tag.Get(string(POST)))
			} else if field.Tag.Get(string(PUT)) != "" {
				routers = append(routers, "PUT", field.Tag.Get(string(PUT)))
			} else if field.Tag.Get(string(PATCH)) != "" {
				routers = append(routers, "PATCH", field.Tag.Get(string(PATCH)))
			} else if field.Tag.Get(string(DELETE)) != "" {
				routers = append(routers, "DELETE", field.Tag.Get(string(DELETE)))
			}

			if len(routers) == 0 {
				panic(errors.New("path register is invalid"))
			}

			// Get handler
			val := ct.Field(i)
			handler := val.Interface().(Handler)

			// register
			endpoint := register(registerOpt{
				method:            routers[0],
				prefix:            IfSlashPrefixString(prefix),
				path:              routers[1],
				GlobalMiddlewares: GlobalMiddlewares,
				middlewares:       middlewares,
				handler:           http.HandlerFunc(handler),
			})

			for k, v := range endpoint {
				mux[k] = v
			}

			// Reset route middlewares
			middlewares = []middleware{}
		}

		// Reset group middlewares
		GlobalMiddlewares = []middleware{}
	}

	return mux
}

type registerOpt struct {
	method            string
	prefix            string
	path              string
	GlobalMiddlewares []middleware
	middlewares       []middleware
	handler           http.Handler
}

func register(opt registerOpt) Mux {
	mux := make(Mux)

	route := fmt.Sprintf("%s %s%s", opt.method, IfSlashPrefixString(opt.prefix), IfSlashPrefixString(opt.path))
	mergeHandler := opt.handler

	for _, m := range opt.GlobalMiddlewares {
		mergeHandler = m(mergeHandler)
	}

	for _, m := range opt.middlewares {
		mergeHandler = m(mergeHandler)
	}

	mux[route] = mergeHandler

	return mux
}
