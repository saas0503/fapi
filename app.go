package api

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Prefix      string
	Middlewares []middleware
	mux         Mux
}

func NewApp(prefix string) *App {
	return &App{
		Prefix:      IfSlashPrefixString(prefix),
		Middlewares: []middleware{},
		mux:         make(Mux),
	}
}

func (a *App) Use(middleware middleware) {
	a.Middlewares = append(a.Middlewares, middleware)
}

func (a *App) Listen(port int) {
	router := http.NewServeMux()

	for k, v := range a.mux {
		router.Handle(k, v)
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
