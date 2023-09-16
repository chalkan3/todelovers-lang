package runtime

import (
	"fmt"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/register"
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
	arg1 := c.Request(func(pm program.ProgramManager) interface{} {
		return pm.GetAdressValue(1, byte(threadID))
	}).ToByte()

	register := c.Request(func(pm register.RegisterManager) interface{} {
		return pm.GetRegister(arg1)
	}).ToRegister()

	fmt.Println(register.Value())

	c.MoveProgramPointer(2, threadID)
}
