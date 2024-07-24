package api

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

type App struct {
	cors        *cors.Cors
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

func (a *App) EnableCors(opt cors.Options) {
	a.cors = cors.New(opt)
}

func (a *App) Listen(port int) {
	router := http.NewServeMux()

	for k, v := range a.mux {
		router.Handle(k, v)
	}

	// Free allocation
	a.mux = nil
	a.Middlewares = nil

	var handler = logRequest(router)
	if a.cors != nil {
		handler = a.cors.Handler(router)
	}
	a.cors = nil

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	log.Printf("Starting server on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("Received request:  %s %s %s", r.Method, r.URL.Path, elapsed)
	})
}
