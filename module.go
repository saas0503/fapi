package api

type Scope string

const (
	GLOBAL  Scope = "global"
	REQUEST Scope = "request"
)

type Module struct {
	Scope Scope
	mux   Mux
}

type ModuleOptions struct {
	Imports    []*Module
	Controller []interface{}
}

func NewModule(opt ModuleOptions) *Module {
	mux := make(Mux)
	for _, m := range opt.Imports {
		for k, v := range m.mux {
			mux[k] = v
		}
		if m.Scope == REQUEST {
			m = nil
		}
	}

	ctrl := Registry(opt.Controller...)
	for k, v := range ctrl {
		mux[k] = v
	}

	return &Module{
		mux:   mux,
		Scope: REQUEST,
	}
}

func CreateApp(module *Module) *App {
	return &App{
		module: module,
	}
}
