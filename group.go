package api

/*
import "github.com/saas0503/factory-api/guard"

type Group struct {
	Name        string
	Middlewares []middleware
	app         *App
}

func (a *App) AddGroup(name string) *Group {
	return &Group{
		Name:        a.Prefix + IfSlashPrefixString(name),
		Middlewares: a.Middlewares,
		app:         a,
	}
}

func (g *Group) Use(middleware middleware) *Group {
	g.Middlewares = append(g.Middlewares, middleware)
	return g
}

func (g *Group) Auth() *Group {
	g.Middlewares = append(g.Middlewares, guard.Authentication)
	return g
}

func (g *Group) Registry() *Router {
	return &Router{
		Name:              g.Name,
		GlobalMiddlewares: g.Middlewares,
		group:             g,
	}
}
*/
