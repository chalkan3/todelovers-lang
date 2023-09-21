package cplaneapi

import (
	"mary_guica/pkg/nando"
	control "mary_guica/pkg/tvm/pkg/control_plane"
)

type RouteKey int

func (r RouteKey) String() string {
	return [...]string{
		"get-program-manager",
		"get-thread-manager",
		"get-register-manager",
		"get-memory-manager",
	}[r]
}

const (
	SERVER = "x-controlplane-api"
)
const (
	Program RouteKey = iota
	Thread
	Register
	Memory
)

type API interface {
	Routes(svc Service) []*nando.Handler
	Serve()
}
type api struct {
	cp control.ControlPlane
}

func NewAPI(cp control.ControlPlane) API {
	return &api{
		cp: cp,
	}
}

func (e *api) Serve() {
	svc := NewService(e.cp)
	go nando.NewServer(SERVER).Listen(e.Routes(svc)...)
}

func (e *api) Routes(svc Service) []*nando.Handler {
	return []*nando.Handler{
		nando.NewHandler(e.cp.ProgramManager().APIPath(), nando.HandleFunc(ProgramManagerEndpoint(svc))),
		nando.NewHandler(e.cp.ThreadManager().APIPath(), nando.HandleFunc(ThreadManagerEndpoint(svc))),
		nando.NewHandler(e.cp.RegisterManager().APIPath(), nando.HandleFunc(RegisterManagerEndpoint(svc))),
		nando.NewHandler(e.cp.MemoryManager().APIPath(), nando.HandleFunc(MemoryManagerEndpoint(svc))),
	}
}

func Client() *nando.Client {
	return &nando.Client{
		Server: SERVER,
	}
}
