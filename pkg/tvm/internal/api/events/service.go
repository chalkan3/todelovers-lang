package events

import (
	wal "mary_guica/pkg/tlwal"
	"mary_guica/pkg/tvm/pkg/events"
	"math/rand"
	"strconv"
)

type Service interface {
	CreateHandler(eventHandler *EventHandler) error
	Notify(notifier *Notifier) error
}

type service struct {
	wal wal.TLWAL
	ec  events.EventController
}

func NewService(ec events.EventController) Service {
	return &service{
		wal: wal.NewTLWAL(),
		ec:  ec,
	}
}

func (s *service) CreateHandler(eventHandler *EventHandler) error {
	record := &wal.Record{
		Operation: "INSERT",
		Table:     "events",
		Data: &wal.Data{
			Key:   eventHandler.ID,
			Value: eventHandler.HandlerName,
		},
	}
	err := s.wal.Write(record, true)
	if err != nil {
		return nil
	}

	s.ec.NewObserver(eventHandler.HandlerName, eventHandler.Handler)
	return nil
}
func (s *service) Notify(notifier *Notifier) error {
	s.ec.Notify(notifier)

	record := &wal.Record{
		Operation: "INSERT",
		Table:     "notifier",
		Data: &wal.Data{
			Key:   strconv.Itoa(rand.Int()),
			Value: notifier.Event.Description,
		},
	}

	err := s.wal.Write(record, true)
	if err != nil {
		return nil
	}
	return nil
}
