package runtime

type store struct {
	*base
}

func NewStore(r FlightAttendant) Command {
	return &store{
		&base{
			requester: r,
		},
	}
}

func (c *store) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	// arg2 := c.GetCurrentThread(threadID).PcPointer(2)

	// registerID := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// saveToAdress := c.GetCurrentThread(threadID).GetMemoryPos(arg2)

	// register := c.GetCurrentThread(threadID).GetRegister(registerID)
	// value, ok := toString(register.Value())

	// if ok {
	// 	for _, char := range value {
	// 		c.GetCurrentThread(threadID).SetMemory(int(saveToAdress), byte(char))
	// 		saveToAdress++
	// 	}
	// } else {
	// 	c.GetCurrentThread(threadID).SetMemory(int(saveToAdress), toByte(register.Value()))
	// }

	// c.GetCurrentThread(threadID).MovePC(3)
}
