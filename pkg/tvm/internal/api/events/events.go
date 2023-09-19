package events

import (
	"mary_guica/pkg/nando"
	wal "mary_guica/pkg/tlwal"
	"mary_guica/pkg/tvm/pkg/events"
	"mary_guica/pkg/tvm/pkg/logger"
	"os"
)

type EventsAPI interface {
	Routes() []*nando.Handler
	Serve()
}
type eventsAPI struct {
	wal wal.TLWAL
	ec  events.EventController
}

type Request struct {
	ID          string
	HandlerName string
}

func NewEventsAPI() EventsAPI {
	return &eventsAPI{
		wal: wal.NewTLWAL(),
		ec:  events.NewEventController(),
	}
}

func (e *eventsAPI) Serve() {
	nando.Listen(e.Routes()...)
}

func (e *eventsAPI) Routes() []*nando.Handler {
	return []*nando.Handler{
		nando.NewHandler("create-handler", func(r *nando.Request) (*nando.Response, error) {
			record := &wal.Record{
				Operation: "INSERT",
				Table:     "events",
				Data: &wal.Data{
					Key:   r.Data.(*Request).ID,
					Value: r.Data.(*Request).HandlerName,
				},
			}
			e.wal.CreateLogFile()
			e.wal.Write(record, true)

			events.GetEventController().NewObserver("NEW_CREW", []events.Observer{
				logger.NewConsoleLogObserver(os.Stdout),
				logger.NewFileLogObserver(),
			})

			events.GetEventController().NewObserver("NOTIFY", []events.Observer{
				logger.NewConsoleLogObserver(os.Stdout),
				logger.NewFileLogObserver(),
			})
			return &nando.Response{}, nil
		}),
		nando.NewHandler("notify", func(*nando.Request) (*nando.Response, error) {

			return &nando.Response{}, nil
		}),
	}
}
