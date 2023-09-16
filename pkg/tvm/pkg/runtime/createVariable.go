package runtime

type createVariable struct {
	*base
}

func NewCreateVariable(r FlightAttendant) Command {
	return &createVariable{
		&base{
			requester: r,
		},
	}
}

func (c *createVariable) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	// nameLength := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// nameBytes := make([]byte, nameLength)

	// for i := byte(0); i < nameLength; i++ {
	// 	argN := int(byte(c.GetCurrentThread(threadID).PcPointer(2)) + i)
	// 	nameBytes[i] = c.GetCurrentThread(threadID).GetMemoryPos(argN)
	// }

	// variableName := string(nameBytes)
	// sizeVariables := c.GetCurrentThread(threadID).LenVariables()

	// c.GetCurrentThread(threadID).CreateVariable(&tvm.VariableParams{
	// 	Name:   variableName,
	// 	Adress: 0xC8 + byte(sizeVariables),
	// 	Value:  0,
	// })

	// c.GetCurrentThread(threadID).SetMemory(0xC8, 0x00)

	// c.GetCurrentThread(threadID).MovePC(int(nameLength) + 2)
}
