package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type App struct {
	Prefix      string
	Middlewares []middleware
	mux         Mux
}

func CreateApp(module *Module) *App {
	app := &App{
		Prefix:      "",
		Middlewares: []middleware{},
		mux:         make(Mux),
	}

	for k, v := range module.mux {
		fmt.Printf("Final path is: %s\n", k)
		app.mux[k] = v
	}
	module = nil

	return app
}

func (a *App) Use(middleware middleware) {
	a.Middlewares = append(a.Middlewares, middleware)
}

func (a *App) SetGlobalPrefix(prefix string) {
	a.Prefix = prefix
}

func (a *App) Listen(port int) {
	router := http.NewServeMux()

	for k, v := range a.mux {
		routes := strings.Split(k, " ")
		path := routes[0] + " " + a.Prefix + routes[1]

		mergeHandler := v
		for _, m := range a.Middlewares {
			mergeHandler = m(mergeHandler)
		}

		router.Handle(path, mergeHandler)
	}

	// Free allocation
	a.mux = nil
	a.Middlewares = nil

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	log.Printf("Starting server on port %d", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
