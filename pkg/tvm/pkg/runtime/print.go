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
	arg1 := getArgument(1, threadID)

	register := registerManager().GetRegister(arg1)

	fmt.Println(register.Value())

	moveProgramPointer(2, threadID)
}
