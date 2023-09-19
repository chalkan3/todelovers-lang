package events

import (
	"mary_guica/pkg/nando"
)

type Endpoint func(req *nando.Request) (*nando.Response, error)

type NotifyRequest struct {
	Notifier *Notifier
}
type NotifyResponse struct{}

func NotifyEndpoint(svc Service) Endpoint {
	return func(req *nando.Request) (*nando.Response, error) {
		bind := toNotifyRequest(req)
		err := svc.Notify(bind.Notifier)
		return &nando.Response{}, err
	}
}

type CreateHandlerRequest struct {
	EventHandler *EventHandler
}
type CreateHandlerResponse struct{}

func CreateHandlerEndpoint(svc Service) Endpoint {
	return func(req *nando.Request) (*nando.Response, error) {
		bind := toCreateHandlerRequest(req)
		err := svc.CreateHandler(bind.EventHandler)
		return &nando.Response{}, err
	}
}
