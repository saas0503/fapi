package api

import (
	"net/http"

	"github.com/saas0503/factory-api/guard"
)

type middleware func(http.Handler) http.Handler

type GroupApi struct {
	Prefix      string
	Middlewares []middleware
}

func (app *App) Group(prefix string) *GroupApi {
	return &GroupApi{
		Prefix:      app.Prefix + prefix,
		Middlewares: app.Middlewares,
	}
}

// Guard Auth
func (g *GroupApi) Auth() *GroupApi {
	g.Middlewares = append(g.Middlewares, guard.Authentication)
	return g
}
