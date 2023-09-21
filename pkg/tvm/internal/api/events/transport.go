package events

import (
	"mary_guica/pkg/nando"
	wal "mary_guica/pkg/tlwal"
	"mary_guica/pkg/tvm/pkg/events"
)

type RouteKey int

func (r RouteKey) String() string {
	return [...]string{"notify", "create-handler"}[r]
}

const (
	SERVER = "x-events-api"
)
const (
	Notify RouteKey = iota
	CreateHandler
)

type API interface {
	Routes(svc Service) []*nando.Handler
	Serve()
}
type api struct {
	wal wal.TLWAL
	ec  events.EventController
}

func NewAPI() API {
	return &api{
		wal: wal.NewTLWAL(),
		ec:  events.NewEventController(),
	}
}

func (e *api) Serve() {
	svc := NewService(e.ec)

	go nando.NewServer(SERVER).Listen(e.Routes(svc)...)

	go e.ec.Listen()
}

func (e *api) Routes(svc Service) []*nando.Handler {
	return []*nando.Handler{
		nando.NewHandler(Notify.String(), nando.HandleFunc(NotifyEndpoint(svc))),
		nando.NewHandler(CreateHandler.String(), nando.HandleFunc(CreateHandlerEndpoint(svc))),
	}
}

func Client() *nando.Client {
	return &nando.Client{
		Server: SERVER,
	}
}
