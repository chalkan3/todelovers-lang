package requester

type Crew interface {
	Register(id int)
	Get(id int) FlightAttendant
	PrepareCrew() <-chan int
}
type crew map[int]FlightAttendant

func NewCrew() Crew {
	return make(crew)
}

func (c crew) Register(id int)            { c[id] = NewFlightAttendant() }
func (c crew) Get(id int) FlightAttendant { return c[id] }

func (c crew) PrepareCrew() <-chan int {
	ch := make(chan int, 1000000)
	for channelID, flightAttendant := range c {
		go func(id int, f FlightAttendant) {
			for {
				<-f.WaitForRequest()
				ch <- id

			}
		}(channelID, flightAttendant)
	}
	return ch
}
