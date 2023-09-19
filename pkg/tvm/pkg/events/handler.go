package events

type Event struct {
	Name        string
	Description string
	Data        interface{}
}

type Observer interface {
	Update(event Event)
}

type Handler interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObserver(event Event)
}

type handler struct {
	observerList []Observer
	name         string
}

func newHandler() Handler {
	return &handler{
		observerList: []Observer{},
	}
}
func (s *handler) AddObserver(observer Observer) {
	s.observerList = append(s.observerList, observer)
}

func (s *handler) RemoveObserver(observer Observer) {
	for i, o := range s.observerList {
		if o == observer {
			s.observerList = append(s.observerList[:i], s.observerList[i+1:]...)
			break
		}
	}
}

func (s *handler) NotifyObserver(event Event) {
	for _, o := range s.observerList {
		o.Update(event)
	}
}
