package runtime

import (
	"fmt"
	"mary_guica/pkg/cast"
)

type add struct {
	*base
}

func NewADD(r FlightAttendant) Command {
	return &add{
		&base{
			requester: r,
		},
	}
}

func (c *add) Execute(instruction byte, threadID int, args ...interface{}) {
	arg1 := cast.ToAlwaysInt(getArgument(1, threadID))
	arg2 := cast.ToAlwaysInt(getArgument(2, threadID))

	fmt.Println(arg1 + arg2)

	moveProgramPointer(1, threadID)

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
