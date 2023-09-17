package runtime

import "log"

type Crew interface {
	Register(id int)
	Get(id int) FlightAttendant
	Crew() map[int]FlightAttendant
	Update()
}
type crew struct {
	c       map[int]FlightAttendant
	runtime Runtime
}

func NewCrew(r Runtime) Crew {
	return &crew{
		c:       make(map[int]FlightAttendant),
		runtime: r,
	}
}

func (c crew) Register(id int)            { c.c[id] = NewFlightAttendant() }
func (c crew) Get(id int) FlightAttendant { return c.c[id] }

func (c crew) Crew() map[int]FlightAttendant { return c.c }

func (c crew) Handler() {
	visitor := NewIsRunningVisitor()
	for id, fa := range c.Crew() {
		fa.Accept(visitor, func(f FlightAttendant) {
			log.Printf("[VM] New Crew ID[%d] \n", id)
			for {
				select {
				case fn := <-fa.WaitForRequest():
					var returnV interface{}
					// if fn, ok := fn.(func(control.ControlPlane) interface{}); ok {
					// 	returnV = fn(c.runtime.ControlPlane())
					// }

					// if fn, ok := fn.(func(events.EventController) interface{}); ok {
					// 	returnV = fn(c.runtime.EventController())
					// }

					if fn, ok := fn.(func(Runtime) interface{}); ok {
						returnV = fn(c.runtime)
					}

					fa.Response(returnV)
				}
			}

		})
	}
}

func (c crew) Update() {
	c.Handler()
}
