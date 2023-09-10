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

func (c *print) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)

	registerID := c.GetVM().GetMemoryPos(arg1)
	register := c.GetVM().GetRegister(registerID)

	fmt.Println(register.Value())

	c.GetVM().MovePC(2)
}
