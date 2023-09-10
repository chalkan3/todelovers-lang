package interpreter

import "mary_guica/pkg/tvm"

type add struct {
	*base
}

func NewADD(vm *tvm.TVM) Command {
	return &add{
		&base{
			tvm: vm,
		},
	}
}

func (c *add) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)
	arg2 := c.GetVM().PcPointer(2)

	register1ID := c.GetVM().GetMemoryPos(arg1)
	register2ID := c.GetVM().GetMemoryPos(arg2)

	reg1 := c.GetVM().GetRegister(register1ID)
	reg2 := c.GetVM().GetRegister(register2ID)
	reg0 := c.GetVM().GetRegister(0x00)

	v1 := toAlwaysInt(reg1.Value())
	v2 := toAlwaysInt(reg2.Value())

	sum := v1 + v2
	reg0.Set(sum)

	c.GetVM().MovePC(3)
}
