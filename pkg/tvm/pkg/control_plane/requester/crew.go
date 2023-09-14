package requester

import (
	"fmt"
	"reflect"
	"time"
)

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

func (c crew) WaitingToServe() {
	crew := c.PrepareCrew()
	for {
		select {
		case channelID := <-crew:
			fn := <-c.Get(channelID).WaitForRequest()
			v := reflect.ValueOf(fn)
			v.Call([]reflect.Value{reflect.ValueOf(rt.cp.MemoryManager())})

		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: No data received in 2 seconds.")
			return
		}
	}
}

func (c crew) PrepareCrew() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for channelID, flightAttendant := range c {
			select {
			case <-flightAttendant.WaitForRequest():
				ch <- channelID
			}
		}
	}()
	return ch
}
