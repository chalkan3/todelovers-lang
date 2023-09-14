package commands

import (
	"fmt"
	control "mary_guica/pkg/tvm/pkg/control_plane/requester"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/threads"
)

type add struct {
	*base
}

func NewADD(r control.FlightAttendant) Command {
	return &add{
		&base{
			requester: r,
		},
	}
}

func (c *add) Execute(instruction byte, threadID int, args ...interface{}) {
	fmt.Println("executei")
	c.Request(func(pm program.ProgramManager) interface{} {
		pm.Next()
		fmt.Println(pm.Instruction())
		return nil
	})

	c.Request(func(pm threads.ThreadManager) interface{} {
		pm.GetThread(threadID).Next()
		return nil
	})
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	// arg2 := c.GetCurrentThread(threadID).PcPointer(2)

	// register1ID := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// register2ID := c.GetCurrentThread(threadID).GetMemoryPos(arg2)

	// reg1 := c.GetCurrentThread(threadID).GetRegister(register1ID)
	// reg2 := c.GetCurrentThread(threadID).GetRegister(register2ID)
	// reg0 := c.GetCurrentThread(threadID).GetRegister(0x00)

	// v1 := toAlwaysInt(reg1.Value())
	// v2 := toAlwaysInt(reg2.Value())

	// sum := v1 + v2
	// reg0.Set(sum)

	// c.GetCurrentThread(threadID).MovePC(3)
}
