package runtime

type FlightAttendant interface {
	Request(fn interface{})
	Response(v interface{})
	WaitForRequest() chan interface{}
	WaitForReponse() chan Output
	Running() bool
	SetRunning()
	Accept(visitor Visitor, fn func(fa FlightAttendant))
}

type flightAttendant struct {
	running  bool
	request  chan interface{}
	executed chan Output
}

func NewFlightAttendant() FlightAttendant {
	return &flightAttendant{
		running:  false,
		request:  make(chan interface{}),
		executed: make(chan Output),
	}
}

func (f *flightAttendant) SetRunning()                      { f.running = true }
func (f *flightAttendant) Running() bool                    { return f.running }
func (f *flightAttendant) WaitForRequest() chan interface{} { return f.request }

func (f *flightAttendant) Request(fn interface{}) {
	f.request <- fn
}

func (f *flightAttendant) WaitForReponse() chan Output {
	return f.executed
}

func (f *flightAttendant) Response(v interface{}) {
	f.executed <- Output{v: v}
}

func (f *flightAttendant) Accept(visitor Visitor, fn func(fa FlightAttendant)) {
	visitor.Visit(f, fn)
}
