package api

type Module struct {
	app *App
	mux Mux
}

type ModuleOptions struct {
	Imports    []*Module
	Controller []*Controller
}

func NewModule(opt ModuleOptions) *Module {
	mux := make(Mux)
	for _, m := range opt.Imports {
		for k, v := range m.mux {
			mux[k] = v
		}
		m = nil
	}

	for _, ctrl := range opt.Controller {
		for k, v := range ctrl.mux {
			mux[k] = v
		}
		ctrl = nil
	}

	return &Module{
		mux: mux,
	}
}

func CreateApp(module *Module) *App {
	return &App{
		mux: module.mux,
	}
}
