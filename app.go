package api

import (
	"fmt"
	"github.com/saas0503/factory-api/interceptor"
	"log"
	"net/http"
)

type App struct {
	Prefix      string
	Middlewares []middleware
	mux         Mux
	routes      map[string]bool
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

func (a *App) routeExists(path string) bool {
	return a.routes[path]
}

func (a *App) Listen(port int) {
	router := http.NewServeMux()
	a.routes = map[string]bool{}

	for k, v := range a.mux {
		fmt.Printf("The path is register %s\n", k)
		router.Handle(k, v)
		a.routes[k] = true
	}

	// Free allocation
	clear(a.mux)
	clear(a.Middlewares)

	var handler = a.routeCheckerMiddleware(router)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	log.Printf("Starting server on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil {

		log.Fatalf("Server failed to start: %s", err)
	}
}

func (a *App) routeCheckerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if route exists
		if a.routeExists(r.Method + " " + r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			interceptor.HandleNotFound(w, r)
			return
		}
	})
}
