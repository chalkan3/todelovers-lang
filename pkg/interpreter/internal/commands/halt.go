package commands

import control "mary_guica/pkg/tvm/pkg/control_plane/requester"

type halt struct {
	*base
}

func NewHalt(r control.FlightAttendant) Command {
	return &halt{
		&base{
			requester: r,
		},
	}
}

func (c *halt) Execute(instruction byte, threadID int, args ...interface{}) {
	// c.GetCurrentThread(threadID).SetDone()
}
