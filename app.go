package fapi

import (
	"fmt"
	"github.com/saas0503/fapi/config"
	"github.com/saas0503/fapi/interceptor"
	"log"
	"net/http"
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

func (a *App) InitConfig(path string) *config.Config {
	cfg, err := config.Load(path)
	if err != nil {
		log.Fatalf("Error when load env: %v", err)
	}
	return cfg
}

func (a *App) Listen(port int) {
	router := http.NewServeMux()
	a.routes = map[string]bool{}

	router.HandleFunc("/", a.ServeHTTP)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	log.Printf("Starting server on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil {

		log.Fatalf("Server failed to start: %s", err)
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.Method + " " + r.URL.Path
	if a.routeExists(path) {
		handler := a.module.mux[path]
		err := handler(w, r)
		if err != nil {
			log.Printf("Error handling request: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	} else {
		interceptor.HandleNotFound(w, r)
		return
	}

}
