package runtime

import (
	"fmt"
)

type print struct {
	*base
}

func NewPrint(r FlightAttendant) Command {
	return &print{
		&base{
			requester: r,
		},
	}
}

func (c *print) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	arg1 := c.GetArgument(1, threadID).ToByte()

	register := c.Request(func(pm Runtime) interface{} {
		return pm.ControlPlane().RegisterManager().GetRegister(arg1)
	}).ToRegister()

	fmt.Println(register.Value())

	c.MoveProgramPointer(2, threadID)
}
