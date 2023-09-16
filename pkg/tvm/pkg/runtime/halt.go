package runtime

import (
	"mary_guica/pkg/tvm/pkg/threads"
)

type halt struct {
	*base
}

func NewHalt(r FlightAttendant) Command {
	return &halt{
		&base{
			requester: r,
		},
	}
}

func (c *halt) Execute(instruction byte, threadID int, args ...interface{}) {
	c.Request(func(pm threads.ThreadManager) interface{} {
		pm.GetThread(threadID).SetDone()
		return nil
	})
}
