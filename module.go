package api

import "net/http"

type Module struct {
	mux Mux
}

type ModuleOptions struct {
	Controllers []*Controller
}

func NewModule(opt ModuleOptions) *Module {
	// Get routes in controllers
	mux := make(map[string]http.Handler)

	for _, c := range opt.Controllers {
		for k, v := range c.mux {
			mux[k] = v
		}
		c = nil
	}

	return &Module{
		mux: mux,
	}
}
