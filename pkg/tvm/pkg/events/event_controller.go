package events

type EventController interface {
	NewEvent(handler string, ev Events)
	NewHandler(name string)
	Notify(handler string)
	Listen()
}
type eventController struct {
	h      map[string]Handler
	notify chan string
}

func (c *eventController) NewHandler(name string) {
	c.h[name] = newHandler()
}

func (c *eventController) NewEvent(handler string, ev Events) {
	c.h[handler].AddEvents(ev)
}

func (c *eventController) Notify(handler string) {
	c.notify <- handler
}

func (c *eventController) Listen() {
	for {
		handler := <-c.notify
		go c.h[handler].NotifyEvents()
	}
}

func NewEventController() EventController {
	return &eventController{
		h:      make(map[string]Handler),
		notify: make(chan string),
	}
}

// controller := NewEventsController()
// controller.Register("")
// controller.Notify("newContext")
// for {}
//
//
