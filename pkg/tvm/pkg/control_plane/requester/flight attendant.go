package requester

type FlightAttendant interface {
	Request(fn interface{})
	Response()
	WaitForRequest() chan interface{}
}
type flightAttendant struct {
	request  chan interface{}
	executed chan bool
}

func NewFlightAttendant() FlightAttendant {
	return &flightAttendant{
		request:  make(chan interface{}),
		executed: make(chan bool),
	}
}

func (f *flightAttendant) WaitForRequest() chan interface{} { return f.request }

func (f *flightAttendant) Request(fn interface{}) {
	f.request <- fn
	// <-f.executed
}

func (f *flightAttendant) Response() {
	f.executed <- true
}
