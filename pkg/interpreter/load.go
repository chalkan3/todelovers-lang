package interpreter

import "mary_guica/pkg/tvm"

type load struct {
	*base
}

func NewLoad(vm *tvm.TVM) Command {
	return &load{
		&base{
			tvm: vm,
		},
	}
}

func (c *load) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)
	arg2 := c.GetVM().PcPointer(2)

	address := c.GetVM().GetMemoryPos(arg1)
	registerID := c.GetVM().GetMemoryPos(arg2)

	register := c.GetVM().GetRegister(registerID)
	register.Set(address)

	c.GetVM().MovePC(3)
}
