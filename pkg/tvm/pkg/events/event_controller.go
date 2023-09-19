package events

type EventController interface {
	NewObserver(name string, os []Observer)
	AddObserver(name string, os []Observer)
	Notify(notifier *Notifier)
	Listen()
}

type Notifier struct {
	Handler string
	Event   *Event
}
type eventController struct {
	h      map[string]Handler
	notify chan *Notifier
}

func (c *eventController) NewObserver(name string, os []Observer) {
	c.h[name] = newHandler()
	c.AddObserver(name, os)

}

func (c *eventController) AddObserver(handler string, os []Observer) {
	for _, o := range os {
		c.h[handler].AddObserver(o)
	}
}

func (c *eventController) Notify(notifier *Notifier) {
	c.notify <- notifier
}

func (c *eventController) Listen() {
	for {
		notifier := <-c.notify
		go c.h[notifier.Handler].NotifyObserver(*notifier.Event)
	}
}

func NewEventController() EventController {
	return &eventController{
		h:      make(map[string]Handler),
		notify: make(chan *Notifier),
	}
}
