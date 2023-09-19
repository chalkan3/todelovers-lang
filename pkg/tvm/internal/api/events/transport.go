package events

import (
	"mary_guica/pkg/nando"
	wal "mary_guica/pkg/tlwal"
	"mary_guica/pkg/tvm/pkg/events"
)

type EventsAPI interface {
	Routes(svc Service) []*nando.Handler
	Serve()
}
type eventsAPI struct {
	wal wal.TLWAL
	ec  events.EventController
}

func NewEventsAPI() EventsAPI {
	return &eventsAPI{
		wal: wal.NewTLWAL(),
		ec:  events.NewEventController(),
	}
}

func (e *eventsAPI) Serve() {
	go e.ec.Listen()
	svc := NewService(e.ec)
	nando.Listen(e.Routes(svc)...)
}

func (e *eventsAPI) Routes(svc Service) []*nando.Handler {
	return []*nando.Handler{
		nando.NewHandler("notify", nando.HandleFunc(NotifyEndpoint(svc))),
		nando.NewHandler("create-handler", nando.HandleFunc(CreateHandlerEndpoint(svc))),
	}
}
