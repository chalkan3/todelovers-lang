package interpreter

import "mary_guica/pkg/tvm"

type store struct {
	*base
}

func NewStore(vm *tvm.TVM) Command {
	return &store{
		&base{
			tvm: vm,
		},
	}
}

func (c *store) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)
	arg2 := c.GetVM().PcPointer(2)

	registerID := c.GetVM().GetMemoryPos(arg1)
	saveToAdress := c.GetVM().GetMemoryPos(arg2)

	register := c.GetVM().GetRegister(registerID)
	value, ok := toString(register.Value())

	if ok {
		for _, char := range value {
			c.GetVM().SetMemory(int(saveToAdress), byte(char))
			saveToAdress++
		}
	} else {
		c.GetVM().SetMemory(int(saveToAdress), toByte(register.Value()))
	}

	c.GetVM().MovePC(3)
}
