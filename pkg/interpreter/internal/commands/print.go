package commands

import control "mary_guica/pkg/tvm/pkg/control_plane/requester"

type print struct {
	*base
}

func NewPrint(r control.FlightAttendant) Command {
	return &print{
		&base{
			requester: r,
		},
	}
}

func (c *print) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)

	// registerID := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// register := c.GetCurrentThread(threadID).GetRegister(registerID)

	// fmt.Println("from thread:", threadID, register.Value())

	// c.GetCurrentThread(threadID).MovePC(2)
}
