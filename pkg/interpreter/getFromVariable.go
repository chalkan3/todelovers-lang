package interpreter

import "mary_guica/pkg/tvm"

type getFromVariable struct {
	*base
}

func NewGetFromVariable(vm *tvm.TVM) Command {
	return &getFromVariable{
		&base{
			tvm: vm,
		},
	}
}

func (c *getFromVariable) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)
	arg2 := c.GetVM().PcPointer(2)

	registerID := c.GetVM().GetMemoryPos(arg1)
	nameLength := c.GetVM().GetMemoryPos(arg2)

	nameBytes := make([]byte, nameLength)

	for i := byte(0); i < nameLength; i++ {
		argN := int(byte(c.GetVM().PcPointer(3)) + i)
		nameBytes[i] = c.GetVM().GetMemoryPos(argN)
	}

	variableName := string(nameBytes)

	variable := c.GetVM().GetVariable(variableName)
	varAdress := variable.GetAdress()

	memoryValue := c.GetVM().GetMemoryPos(int(varAdress))
	register := c.GetVM().GetRegister(registerID)
	register.Set(memoryValue)

	c.GetVM().MovePC(int(nameLength) + 3)
}
