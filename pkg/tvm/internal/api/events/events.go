package events

import (
	"mary_guica/pkg/nando"
	wal "mary_guica/pkg/tlwal"
)

type EventsAPI interface {
	Routes() []*nando.Handler
}
type eventsAPI struct {
	wal wal.TLWAL
}

func NewEventsAPI() EventsAPI {
	return &eventsAPI{
		wal: wal.NewTLWAL(),
	}
}

func (e *eventsAPI) Routes() []*nando.Handler {
	return []*nando.Handler{
		nando.NewHandler("create-handler", func(*nando.Request) (*nando.Response, error) {
			record := wal.WALRecord{
				Operation: "INSERT",
				Data: map[string]interface{}{
					"key":   "1",
					"value": "Hello, World!",
				},
			}
			e.wal.CreateLogFile()
			e.wal.Write(record)
			return &nando.Response{}, nil
		}),
		nando.NewHandler("notify", func(*nando.Request) (*nando.Response, error) {

			return &nando.Response{}, nil
		}),
	}
}
