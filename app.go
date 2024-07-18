package api

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Prefix string
	mux    Mux
}

func Create(prefix string) *App {
	return &App{
		Prefix: prefix,
	}
}

func (a *App) Registry(name string, module *Module) {
	for k, v := range module.mux {
		path := a.Prefix + name + k
		a.mux[path] = v
	}
	module.mux = nil
}

func (a *App) Listen(port int) {
	router := http.NewServeMux()

	for k, v := range a.mux {
		router.Handle(k, v)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	log.Printf("Starting server on port %d", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}