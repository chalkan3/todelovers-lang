package events

import "mary_guica/pkg/tvm/pkg/events"

type Notifier = events.Notifier

type EventHandler struct {
	ID          string
	HandlerName string
	Handler     []events.Observer
}
