package api

type Module struct {
	Imports    []*Module
	Controller []*Controller
}
