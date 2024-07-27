package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/saas0503/factory-api/interceptor"
)

type App struct {
	Prefix string
	module *Module
	routes map[string]bool
}

func (a *App) routeExists(path string) bool {
	return a.routes[path]
}

func (a *App) SetGlobalPrefix(prefix string) *App {
	a.Prefix = IfSlashPrefixString(prefix)
	return a
}

func (a *App) Listen(port int) {
	router := http.NewServeMux()
	a.routes = map[string]bool{}

	for k, v := range a.module.mux {
		routes := strings.Split(k, " ")
		path := routes[0] + " " + a.Prefix + routes[1]
		fmt.Printf("The path is register %s\n", path)
		router.Handle(path, v)
		a.routes[path] = true
	}

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
