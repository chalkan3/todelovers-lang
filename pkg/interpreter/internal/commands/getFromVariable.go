package commands

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

func (c *getFromVariable) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	// arg2 := c.GetCurrentThread(threadID).PcPointer(2)

	// registerID := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// nameLength := c.GetCurrentThread(threadID).GetMemoryPos(arg2)

	// nameBytes := make([]byte, nameLength)

	// for i := byte(0); i < nameLength; i++ {
	// 	argN := int(byte(c.GetCurrentThread(threadID).PcPointer(3)) + i)
	// 	nameBytes[i] = c.GetCurrentThread(threadID).GetMemoryPos(argN)
	// }

	// variableName := string(nameBytes)

	// variable := c.GetCurrentThread(threadID).GetVariable(variableName)
	// varAdress := variable.GetAdress()

	// memoryValue := c.GetCurrentThread(threadID).GetMemoryPos(int(varAdress))
	// register := c.GetCurrentThread(threadID).GetRegister(registerID)
	// register.Set(memoryValue)

	// c.GetCurrentThread(threadID).MovePC(int(nameLength) + 3)
}
