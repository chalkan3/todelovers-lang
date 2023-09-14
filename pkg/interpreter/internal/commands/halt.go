package commands

import (
	"fmt"
	control "mary_guica/pkg/tvm/pkg/control_plane/requester"
	"mary_guica/pkg/tvm/pkg/program"
)

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
	fmt.Println("executei")
	c.Request(func(pm program.ProgramManager) interface{} {
		fmt.Println(pm.Instruction())
		return nil
	})
	// c.GetCurrentThread(threadID).SetDone()
}
