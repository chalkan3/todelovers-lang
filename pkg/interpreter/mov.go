package interpreter

import "mary_guica/pkg/tvm"

type mov struct {
	*base
}

func NewMov(vm *tvm.TVM) Command {
	return &mov{
		&base{
			tvm: vm,
		},
	}
}

func (c *mov) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)
	arg2 := c.GetVM().PcPointer(2)

	fromRegisterID := c.GetVM().GetMemoryPos(arg1)
	toRegisterID := c.GetVM().GetMemoryPos(arg2)

	fromReg := c.GetVM().GetRegister(fromRegisterID)
	toReg := c.GetVM().GetRegister(toRegisterID)

	toReg.Set(fromReg.Value())
	fromReg.Set(1)

	c.GetVM().MovePC(3)
}
