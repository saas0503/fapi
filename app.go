package api

type App struct {
	Port        uint
	Prefix      string
	Middlewares []middleware
	Groups      []GroupApi
}

func New() {

}

func (a *App) SetGlobalMiddleware(m ...middleware) *App {
	a.Middlewares = append(a.Middlewares, m...)

	return a
}
