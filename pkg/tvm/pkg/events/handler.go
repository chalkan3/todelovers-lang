package events

type Events interface {
	Update()
}

type Handler interface {
	AddEvents(observer Events)
	RemoveEvents(observer Events)
	NotifyEvents()
}

type handler struct {
	observerList []Events
	name         string
}

func newHandler() Handler {
	return &handler{
		observerList: []Events{},
	}
}
func (s *handler) AddEvents(observer Events) {
	s.observerList = append(s.observerList, observer)
}

func (s *handler) RemoveEvents(observer Events) {
	for i, o := range s.observerList {
		if o == observer {
			s.observerList = append(s.observerList[:i], s.observerList[i+1:]...)
			break
		}
	}
}

func (s *handler) NotifyEvents() {
	for _, o := range s.observerList {
		o.Update()
	}
}
