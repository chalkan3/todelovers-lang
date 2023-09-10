package interpreter

import "mary_guica/pkg/tvm"

type loadString struct {
	*base
}

func NewLoadString(vm *tvm.TVM) Command {
	return &loadString{
		&base{
			tvm: vm,
		},
	}
}

func (c *loadString) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)
	arg2 := c.GetVM().PcPointer(2)

	registerID := c.GetVM().GetMemoryPos(arg1)
	register := c.GetVM().GetRegister(registerID)

	length := c.GetVM().GetMemoryPos(arg2)

	strData := make([]byte, length)

	for i := byte(0); i < length; i++ {
		n := byte(c.GetVM().PcPointer(3)) + i
		strData[i] = c.GetVM().GetMemoryPos(int(n))
	}

	str := string(strData)

	register.Set(str)

	c.GetVM().MovePC(3 + int(length))
}
