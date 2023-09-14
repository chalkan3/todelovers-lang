package requester

type FlightAttendant interface {
	Request(fn interface{})
	WaitForRequest() chan interface{}
}
type flightAttendant struct {
	request chan interface{}
}

func NewFlightAttendant() FlightAttendant {
	return &flightAttendant{
		request: make(chan interface{}),
	}
}

func (f *flightAttendant) WaitForRequest() chan interface{} { return f.request }

func (f *flightAttendant) Request(fn interface{}) {
	f.request <- fn
}
