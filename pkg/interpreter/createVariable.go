package interpreter

import "mary_guica/pkg/tvm"

type createVariable struct {
	*base
}

func NewCreateVariable(vm *tvm.TVM) Command {
	return &createVariable{
		&base{
			tvm: vm,
		},
	}
}

func (c *createVariable) Execute(instruction byte) {
	arg1 := c.GetVM().PcPointer(1)
	nameLength := c.GetVM().GetMemoryPos(arg1)
	nameBytes := make([]byte, nameLength)

	for i := byte(0); i < nameLength; i++ {
		argN := int(byte(c.GetVM().PcPointer(2)) + i)
		nameBytes[i] = c.GetVM().GetMemoryPos(argN)
	}

	variableName := string(nameBytes)
	sizeVariables := c.GetVM().LenVariables()

	c.GetVM().CreateVariable(&tvm.VariableParams{
		Name:   variableName,
		Adress: 0xC8 + byte(sizeVariables),
		Value:  0,
	})

	c.GetVM().SetMemory(0xC8, 0x00)

	c.GetVM().MovePC(int(nameLength) + 2)
}
