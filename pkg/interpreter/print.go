package interpreter

import (
	"fmt"
	"mary_guica/pkg/tvm"
)

type print struct {
	*base
}

func NewPrint(vm *tvm.TVM) Command {
	return &print{
		&base{
			tvm: vm,
		},
	}
}

func (c *print) Execute(instruction byte, threadID int, args ...interface{}) {
	arg1 := c.GetCurrentThread(threadID).PcPointer(1)

	registerID := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	register := c.GetCurrentThread(threadID).GetRegister(registerID)

	fmt.Println(register.Value())

	c.GetCurrentThread(threadID).MovePC(2)
}
